package domain

import (
	"github.com/c12s/oort/poller/domain/async"
	"github.com/c12s/oort/poller/domain/store/outbox"
	"time"
)

type Poller struct {
	store     outbox.Store
	publisher async.Publisher
	done      chan bool
}

func NewPoller(store outbox.Store, publisher async.Publisher) Poller {
	return Poller{
		store:     store,
		publisher: publisher,
	}
}

//TODO: DOCUMENT

func (p Poller) Start(intervalInMs int) {
	p.done = make(chan bool)
	ticker := time.NewTicker(time.Duration(intervalInMs) * time.Millisecond)

	for {
		select {
		case <-p.done:
			ticker.Stop()
			return
		case <-ticker.C:
			go func() {
				messages, err := p.store.GetUnprocessed()
				if err != nil {
					return
				}
				for _, message := range messages {
					err = p.publisher.Publish(message.Kind, message.Payload)
					if err != nil {
						break
					}
					err = p.store.DeleteById(message)
				}
			}()
		}
	}
}

func (p Poller) Stop() {
	if p.done == nil {
		return
	}
	p.done <- true
}
