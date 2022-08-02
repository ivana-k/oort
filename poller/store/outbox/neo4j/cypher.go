package neo4j

import "github.com/c12s/oort/poller/domain/model"

func reserveAndGetUnprocessedCypher() (string, map[string]interface{}) {
	return "MATCH (msg:OutboxMessage{processing: false}) " +
			"SET msg.processing = true " +
			"RETURN id(msg), msg.kind, msg.payload",
		map[string]interface{}{}
}

func setUnprocessedCypher(message model.OutboxMessage) (string, map[string]interface{}) {
	return "MATCH (msg:OutboxMessage{}) " +
			"WHERE id(msg) = $id " +
			"SET msg.processing = false",
		map[string]interface{}{"id": message.Id}
}

func deleteByIdCypher(message model.OutboxMessage) (string, map[string]interface{}) {
	return "DELETE (msg:OutboxMessage{}) " +
			"WHERE id(msg) = $id",
		map[string]interface{}{"id": message.Id}
}
