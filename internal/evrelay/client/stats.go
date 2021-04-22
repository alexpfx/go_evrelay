package client

import (
	"context"
	"fmt"
	"github.com/alexpfx/go_evrelay/internal/evdev"
	"github.com/alexpfx/go_evrelay/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"sort"
	"strings"
)

type Stats interface {
	Start()
	Stop()
	Show()
}

type stats struct {
	port    int
	host    string
	counter map[uint32]int
	keymap  evdev.KeyMap
}

func (s *stats) Start() {
	s.keymap = evdev.NewKeyMap()

	url := fmt.Sprintf("%s:%d", s.host, s.port)

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("não pode conectar ao serviço: %s %s", url, err.Error())
	}

	defer conn.Close()

	client := pb.NewInputEventClient(conn)

	stream, err := client.GetEvents(context.Background(), keypressFilter())
	if err != nil {
		log.Fatal(err)
	}

	s.recvLoop(stream)
}

func (s *stats) recvLoop(stream pb.InputEvent_GetEventsClient) {
	for {
		keyEv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		key := keyEv.Key

		if _, ok := s.counter[key]; ok {
			s.counter[key] += 1
		} else {
			s.counter[key] = 1
		}
	}
}

func keypressFilter() *pb.Filter {
	return &pb.Filter{
		AcceptEvents: []int32{int32(evdev.KeyPress)},
	}
}
func (s stats) count() int {
	c := 0
	for _, v := range s.counter {
		c += v
	}
	return c
}

func (s stats) Stop() {


}

func (s stats) Show() {
	sb := strings.Builder{}
	sorted := make([]countKey, 0)
	counter := s.counter

	for k, v := range counter {
		sorted = append(sorted, countKey{count: v, key: k})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count < sorted[j].count
	})

	for i := range sorted {
		kv := sorted[i]
		sb.WriteString(fmt.Sprintf("0x%-3x %-10s %-5d\n",
			kv.key, s.keymap.GetName(uint16(kv.key)), kv.count))
	}
	fmt.Println(sb.String())
}

func NewStat(host string, port int) Stats {
	return &stats{
		port:    port,
		host:    host,
		counter: map[uint32]int{},
	}
}

type countKey struct {
	count int
	key   uint32
}
