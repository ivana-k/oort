package api

import (
	"errors"
	"github.com/c12s/magnetar/pkg/messaging"
	"log"
)

type AsyncAdministratorClient struct {
	publisher         messaging.Publisher
	subscriberFactory func(subject string) messaging.Subscriber
}

func NewAsyncAdministratorClient(publisher messaging.Publisher, subscriberFactory func(subject string) messaging.Subscriber) (*AsyncAdministratorClient, error) {
	if publisher == nil {
		return nil, errors.New("publisher is nil")
	}
	return &AsyncAdministratorClient{
		publisher:         publisher,
		subscriberFactory: subscriberFactory,
	}, nil
}

func (n *AsyncAdministratorClient) SendRequest(req AdministrationReq, callback AdministrationCallback) error {
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
