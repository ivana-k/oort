package api

import (
	"fmt"
	"github.com/c12s/magnetar/pkg/messaging"
	"github.com/c12s/magnetar/pkg/messaging/nats"
	natsgo "github.com/nats-io/nats.go"
	"log"
)

type AdministrationAsyncClient struct {
	publisher         messaging.Publisher
	subscriberFactory func(subject string) messaging.Subscriber
}

func NewAdministrationAsyncClient(natsAddress string) (*AdministrationAsyncClient, error) {
	conn, err := natsgo.Connect(fmt.Sprintf("nats://%s", natsAddress))
	if err != nil {
		return nil, err
	}
	publisher, err := nats.NewPublisher(conn)
	if err != nil {
		return nil, err
	}
	subscriberFactory := func(subject string) messaging.Subscriber {
		subscriber, _ := nats.NewSubscriber(conn, subject, "")
		return subscriber
	}
	return &AdministrationAsyncClient{
		publisher:         publisher,
		subscriberFactory: subscriberFactory,
	}, nil
}

func (n *AdministrationAsyncClient) SendRequest(req AdministrationReq, callback AdministrationCallback) error {
	reqMarshalled, err := req.Marshal()
	if err != nil {
		return err
	}
	adminReq := &AdministrationAsyncReq{
		Kind:          req.Kind(),
		ReqMarshalled: reqMarshalled,
	}
	adminReqMarshalled, err := adminReq.Marshal()
	if err != nil {
		return err
	}

	// handle responses
	replySubject := n.publisher.GenerateReplySubject()
	subscriber := n.subscriberFactory(replySubject)
	err = subscriber.Subscribe(func(msg []byte, _ string) {
		resp := &AdministrationAsyncResp{}
		err := resp.Unmarshal(msg)
		if err != nil {
			log.Println(err)
			return
		}
		callback(resp)
	})
	if err != nil {
		return err
	}

	// send request
	err = n.publisher.Request(adminReqMarshalled, AdministrationReqSubject, replySubject)
	if err != nil {
		_ = subscriber.Unsubscribe()
		return err
	}
	return nil
}

type AdministrationCallback func(resp *AdministrationAsyncResp)
