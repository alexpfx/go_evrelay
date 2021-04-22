package evrelay

import (
	"fmt"
	"github.com/alexpfx/go_evrelay/internal/evdev"
	"github.com/alexpfx/go_evrelay/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

type server struct {
	address string
	port    int
}

func (s *server) GetEvents(filter *pb.Filter, stream pb.InputEvent_GetEventsServer) error {
	kbs := evdev.GetKeyboards()

	if len(kbs) < 1 {
		return fmt.Errorf("nenhum dispositivo encontrado")
	}

	if p, ok := peer.FromContext(stream.Context()); ok {
		log.Println("cliente conectado: ", p.Addr)
	}

	chs := make([]chan evdev.KeyEvent, 0)
	for _, kb := range kbs {
		c, _ := kb.ReadKeys()
		chs = append(chs, c)
	}
	ch := mergeChans(chs...)

	for ev := range ch {
		if !isAccepted(filter.AcceptEvents, ev.Type) {
			continue
		}

		err := stream.Send(&pb.KeyEvent{
			Time:   ev.Time.Unix(),
			Key:    uint32(ev.Key),
			Type:   int32(ev.Type),
			KeyStr: "",
		})

		if err != nil {
			break
		}
	}

	return nil
}

type Server interface {
	Start()
}

func NewServer(address string, port int) Server {
	return &server{
		address: address,
		port:    port,
	}
}

func (s *server) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.address, s.port))
	if err != nil {
		log.Fatalf("Não pode iniciar o servidor [%s:%d]: %v",
			s.address, s.port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInputEventServer(grpcServer, &server{})


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Não pode iniciar o servidor [%s:%d]: %v",
			s.address, s.port, err)
	}

}

func mergeChans(chs ...chan evdev.KeyEvent) chan evdev.KeyEvent {
	merged := make(chan evdev.KeyEvent)

	for _, c := range chs {
		go func() {
			for v := range c {
				merged <- v
			}
		}()
	}

	return merged
}

func isAccepted(acceptTypes []int32,
	evType evdev.KeyEventType) bool {

	if len(acceptTypes) == 0 {
		return false
	}
	for _, t := range acceptTypes {
		if int32(evType) == t {
			return true
		}

	}
	return false

}
