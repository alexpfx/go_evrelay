package evrelay

import "github.com/alexpfx/go_evrelay/internal/evdev"

type server struct {
}

func (s server) GetEvents(filter *pb.Filter, stream pb.InputEvent_GetEventsServer) error {
	kbs := device.GetKeyboards()
	if len(kbs) < 1 {
		return fmt.Errorf("Nenhum dispositivo encontrado")
	}

	if p, ok := peer.FromContext(stream.Context()); ok {
		log.Println("cliente conectado: ", p.Addr)

	}

	ch, qch := kb.ReadKeys()

	for ev := range ch {
		if !isAccept(filter.AcceptEvents, ev.Type) {
			continue
		}

		err := stream.Send(&pb.KeyEvent{
			Time:   ev.Time.Unix(),
			Key:    uint32(ev.Key),
			Type:   int32(ev.Type),
			KeyStr: device.KeyMap[ev.Key],
		})
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	qch <- true

	return nil
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
