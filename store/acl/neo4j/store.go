package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/model/checker"
	"github.com/c12s/oort/domain/model/syncer"
	"github.com/c12s/oort/domain/store"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AclStore struct {
	manager *TransactionManager
}

func NewAclStore(manager *TransactionManager) store.AclStore {
	return AclStore{
		manager: manager,
	}
}

func (store AclStore) ConnectResources(req syncer.ConnectResourcesReq) syncer.ConnectResourcesResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(connectResourcesCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer.ConnectResourcesResp{
		Resp: syncer.SyncResp{Error: err},
	}
}

func (store AclStore) DisconnectResources(req syncer.DisconnectResourcesReq) syncer.DisconnectResourcesResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(disconnectResourcesCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer.DisconnectResourcesResp{
		Resp: syncer.SyncResp{Error: err},
	}
}

func (store AclStore) UpsertAttribute(req syncer.UpsertAttributeReq) syncer.UpsertAttributeResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(upsertAttributeCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer.UpsertAttributeResp{
		Resp: syncer.SyncResp{Error: err},
	}
}

func (store AclStore) RemoveAttribute(req syncer.RemoveAttributeReq) syncer.RemoveAttributeResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(removeAttributeCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer.RemoveAttributeResp{
		Resp: syncer.SyncResp{Error: err},
	}
}

func (store AclStore) GetAttributes(req checker.GetAttributeReq) checker.GetAttributeResp {
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
		return checker.GetAttributeResp{Attributes: nil, Error: err}
	}
	return checker.GetAttributeResp{Attributes: getAttributes(results), Error: nil}
}

func (store AclStore) InsertPermission(req syncer.InsertPermissionReq) syncer.InsertPermissionResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(insertPermissionCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer.InsertPermissionResp{
		Resp: syncer.SyncResp{Error: err},
	}
}

func (store AclStore) RemovePermission(req syncer.RemovePermissionReq) syncer.RemovePermissionResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(removePermissionCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer.RemovePermissionResp{
		Resp: syncer.SyncResp{Error: err},
	}
}

func (store AclStore) GetPermissionByPrecedence(req checker.GetPermissionReq) checker.GetPermissionByPrecedenceResp {
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
		return checker.GetPermissionByPrecedenceResp{Hierarchy: nil, Error: err}
	}

	distMap := make(map[int]model.PermissionList)
	for _, result := range results.([]interface{}) {
		distance := int(result.([]interface{})[1].(int64))
		_, ok := distMap[distance]
		if !ok {
			distMap[distance] = make([]model.Permission, 0)
		}
		distMap[distance] = append(distMap[distance], getPermission(result.([]interface{})[0]))
	}

	return checker.GetPermissionByPrecedenceResp{Hierarchy: sortByDistanceAsc(distMap), Error: nil}
}
