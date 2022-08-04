package poller

import (
	"fmt"
	"github.com/c12s/oort/poller/domain/async"
	"github.com/c12s/oort/poller/domain/store/outbox"
	"time"
)

type Handler struct {
	store     outbox.Store
	publisher async.Publisher
	done      chan bool
}

func New(store outbox.Store, publisher async.Publisher) Handler {
	return Handler{
		store:     store,
		publisher: publisher,
	}
}

func (h Handler) Start(intervalInMs int) {
	h.done = make(chan bool)
	ticker := time.NewTicker(time.Duration(intervalInMs) * time.Millisecond)

	for {
		select {
		case <-h.done:
			ticker.Stop()
			return
		case <-ticker.C:
			messages, err := h.store.ReserveAndGetUnprocessed()
			if err != nil {
				break
			}
			for _, message := range messages {
				err = h.publisher.Publish(message.Kind, message.Payload)
				if err != nil {
					fmt.Println(err)
					err = h.store.SetUnprocessed(message)
					fmt.Println(err)
					break
				}
				err = h.store.DeleteById(message)
				fmt.Println(err)
			}
		}
	}
}

func (h Handler) Stop() {
	if h.done == nil {
		return
	}
	h.done <- true
}
