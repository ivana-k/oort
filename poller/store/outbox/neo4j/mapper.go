package neo4j

import (
	"fmt"
	"github.com/c12s/oort/poller/domain/model"
)

func getOutboxMessages(cypherResults interface{}) []model.OutboxMessage {
	messages := make([]model.OutboxMessage, 0)
	fmt.Println(cypherResults)
	for _, result := range cypherResults.([]interface{}) {
		messageAttrs := result.([]interface{})
		message := model.OutboxMessage{
			Id:      messageAttrs[0].(int64),
			Kind:    messageAttrs[1].(string),
			Payload: messageAttrs[2].([]byte),
		}
		messages = append(messages, message)
	}
	return messages
}
