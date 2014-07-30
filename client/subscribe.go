package client

import (
	proto "github.com/bjeanes/go-lifx/protocol"
)

func (c *client) Sub() <-chan proto.Message {
	ch := make(chan proto.Message)
	// Doubt this append and replace is atomic which means Sub()s from multiple
	// goroutines could result in lost subscribers
	c.subs.subscribers = append(c.subs.subscribers, subscriber{dst: ch})
	return ch
}

type subscriber struct {
	dst      chan<- proto.Message
	critical bool
}
type subService struct {
	source      <-chan proto.Message
	subscribers []subscriber
}

func newSubService(msgs <-chan proto.Message) (subService, <-chan proto.Message) {
	s := subService{source: msgs}
	clone := make(chan proto.Message)
	s.subscribers = append(s.subscribers, subscriber{clone, true})
	go s.start()
	return s, clone
}

func (s *subService) start() {
	for msg := range s.source {
		for _, sub := range s.subscribers {
			if sub.critical || true {
				// bro, you better not block.
				sub.dst <- msg
			} else {
				select {
				case sub.dst <- msg:
					// default:
					// if a subscriber is lagging, then their messages will be dropped
				}
			}
		}
	}
}
