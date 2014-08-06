package client

import (
	"fmt"
	proto "github.com/bjeanes/go-lifx/protocol"
	"reflect"
)

type subscriber struct {
	dst      chan<- proto.Message
	critical bool
}
type subService struct {
	source      <-chan proto.Message
	subscribing chan chan proto.Message
}

func newSubService(msgs <-chan proto.Message) subService {
	s := subService{source: msgs, subscribing: make(chan chan proto.Message)}
	go s.start()
	return s
}

func (s *subService) Sub() <-chan proto.Message {
	ch := make(chan proto.Message)
	s.subscribing <- ch
	return ch
}

func (s *subService) start() {
	subs := make([]chan<- proto.Message, 0)
	cases := make([]reflect.SelectCase, 0)

	go func() {
		for {
			select {
			// Grow subscriber channels and select cases for each new subscriber
			case ch := <-s.subscribing:
				subs = append(subs, ch)
				cases = append(cases, reflect.SelectCase{Dir: reflect.SelectSend})
			}
		}
	}()

	for msg := range s.source {
		fmt.Printf("%d %+v\n", len(subs), msg)

		// Initialize all cases with the correct channel and message value
		for i, sub := range cases {
			if sub.Dir != reflect.SelectDefault {
				fmt.Printf("preparing case for %v\n", subs[i])
				sub.Chan = reflect.ValueOf(subs[i])
				sub.Send = reflect.ValueOf(msg)
			}
		}

		// For all cases, send a message to the first available one, then remove it
		// and repeat until all cases have been processed.
		for _ = range cases {
			fmt.Println("trying case")
			chosen, _, _ := reflect.Select(cases)
			cases[chosen].Chan = reflect.ValueOf(nil)
		}
	}
}
