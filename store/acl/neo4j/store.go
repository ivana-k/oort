package neo4j

import (
	"errors"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AclStore struct {
	manager *TransactionManager
}

func NewAclStore(manager *TransactionManager) acl.Store {
	return AclStore{
		manager: manager,
	}
}

func (store AclStore) ConnectResources(req acl.ConnectResourcesReq) acl.SyncResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(connectResourcesCypher(req))
		outboxMessage := req.Callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return acl.SyncResp{Error: err}
}

func (store AclStore) DisconnectResources(req acl.DisconnectResourcesReq) acl.SyncResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(disconnectResourcesCypher(req))
		outboxMessage := req.Callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return acl.SyncResp{Error: err}
}

func (store AclStore) UpsertAttribute(req acl.UpsertAttributeReq) acl.SyncResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(upsertAttributeCypher(req))
		outboxMessage := req.Callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return acl.SyncResp{Error: err}
}

func (store AclStore) RemoveAttribute(req acl.RemoveAttributeReq) acl.SyncResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(removeAttributeCypher(req))
		outboxMessage := req.Callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetAttributes(req acl.GetAttributeReq) acl.GetAttributeResp {
	results, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(getAttributeCypher(req))
		if err != nil {
			return nil, err
		}
		if result.Err() != nil {
			return nil, result.Err()
		}

		records := make([]interface{}, 0)
		for result.Next() {
			records = append(records, result.Record().Values[0])
		}
		return records, nil
	})

	if err != nil {
		return acl.GetAttributeResp{Attributes: nil, Error: err}
	}
	return acl.GetAttributeResp{Attributes: getAttributes(results), Error: nil}
}

func (store AclStore) InsertPermission(req acl.InsertPermissionReq) acl.SyncResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(insertPermissionCypher(req))
		outboxMessage := req.Callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return acl.SyncResp{Error: err}
}

func (store AclStore) RemovePermission(req acl.RemovePermissionReq) acl.SyncResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(removePermissionCypher(req))
		outboxMessage := req.Callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetPermissionByPrecedence(req acl.GetPermissionReq) acl.GetPermissionByPrecedenceResp {
	results, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(getPermissionAndDistanceToPrincipal(req))
		if err != nil {
			return nil, err
		}
		if result.Err() != nil {
			return nil, result.Err()
		}

		records := make([]interface{}, 0)
		for result.Next() {
			records = append(records, result.Record().Values)
		}
		return records, nil
	})

	if err != nil {
		return acl.GetPermissionByPrecedenceResp{Hierarchy: nil, Error: err}
	}

	distMap := make(map[int]model.PermissionList)
	for _, result := range results.([]interface{}) {
		distance := int(result.([]interface{})[1].(int64))
		_, ok := distMap[distance]
		if !ok {
			distMap[distance] = make([]model.Permission, 0)
		}
		permission, err := getPermission(result.([]interface{})[0])
		if err != nil {
			return acl.GetPermissionByPrecedenceResp{Hierarchy: nil, Error: err}
		}
		distMap[distance] = append(distMap[distance], permission)
	}

	return acl.GetPermissionByPrecedenceResp{Hierarchy: sortByDistanceAsc(distMap), Error: nil}
}
