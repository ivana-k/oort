package neo4j

import (
	"github.com/c12s/oort/domain/model"
	checker2 "github.com/c12s/oort/domain/model/checker"
	syncer2 "github.com/c12s/oort/domain/model/syncer"
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

func (store AclStore) ConnectResources(req syncer2.ConnectResourcesReq) syncer2.ConnectResourcesResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(connectResourcesCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer2.ConnectResourcesResp{
		Resp: syncer2.SyncResp{Error: err},
	}
}

func (store AclStore) DisconnectResources(req syncer2.DisconnectResourcesReq) syncer2.DisconnectResourcesResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(disconnectResourcesCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer2.DisconnectResourcesResp{
		Resp: syncer2.SyncResp{Error: err},
	}
}

func (store AclStore) UpsertAttribute(req syncer2.UpsertAttributeReq) syncer2.UpsertAttributeResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(upsertAttributeCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer2.UpsertAttributeResp{
		Resp: syncer2.SyncResp{Error: err},
	}
}

func (store AclStore) RemoveAttribute(req syncer2.RemoveAttributeReq) syncer2.RemoveAttributeResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(removeAttributeCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer2.RemoveAttributeResp{
		Resp: syncer2.SyncResp{Error: err},
	}
}

func (store AclStore) GetAttributes(req checker2.GetAttributeReq) checker2.GetAttributeResp {
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
		return checker2.GetAttributeResp{Attributes: nil, Error: err}
	}
	return checker2.GetAttributeResp{Attributes: getAttributes(results), Error: nil}
}

func (store AclStore) InsertPermission(req syncer2.InsertPermissionReq) syncer2.InsertPermissionResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(insertPermissionCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer2.InsertPermissionResp{
		Resp: syncer2.SyncResp{Error: err},
	}
}

func (store AclStore) RemovePermission(req syncer2.RemovePermissionReq) syncer2.RemovePermissionResp {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(removePermissionCypher(req))
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return syncer2.RemovePermissionResp{
		Resp: syncer2.SyncResp{Error: err},
	}
}

func (store AclStore) GetPermissionByPrecedence(req checker2.GetPermissionReq) checker2.GetPermissionByPrecedenceResp {
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
		return checker2.GetPermissionByPrecedenceResp{Hierarchy: nil, Error: err}
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

	return checker2.GetPermissionByPrecedenceResp{Hierarchy: sortByDistanceAsc(distMap), Error: nil}
}
