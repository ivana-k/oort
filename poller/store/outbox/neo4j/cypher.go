package neo4j

import "github.com/c12s/oort/poller/domain/model"

func getUnprocessedCypher() (string, map[string]interface{}) {
	return "MATCH (msg:OutboxMessage{}) " +
			"RETURN id(msg), msg.kind, msg.payload",
		map[string]interface{}{}
}

func deleteByIdCypher(message model.OutboxMessage) (string, map[string]interface{}) {
	return "MATCH (msg:OutboxMessage{}) " +
			"WHERE id(msg) = $id " +
			"DELETE msg",
		map[string]interface{}{"id": message.Id}
}
