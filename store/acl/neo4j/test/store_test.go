package test

import (
	"context"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
	"github.com/c12s/oort/domain/syncer"
	"github.com/c12s/oort/proto/syncerpb"
	neo4jstore "github.com/c12s/oort/store/acl/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	aclStore  acl.Store
	txManager *neo4jstore.TransactionManager
)

func TestResourceConnections(t *testing.T) {
	setUpAclStore(t)

	org := model.NewResource("org1", "org")
	group := model.NewResource("g1", "group")
	user := model.NewResource("u1", "user")

	t.Run("Successfully connect nonexistent resources", func(t *testing.T) {
		parent := org
		child := group

		parentStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Resource)
		assert.ErrorIs(t, parentStored.Error, acl.ErrNotFound)
		childStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Resource)
		assert.ErrorIs(t, childStored.Error, acl.ErrNotFound)

		req := syncer.ConnectResourcesReq{
			ReqId:  "1",
			Parent: parent,
			Child:  child,
		}
		resp := aclStore.ConnectResources(acl.ConnectResourcesReq{
			Parent:   req.Parent,
			Child:    req.Child,
			Callback: outboxMessageCallback(req),
		})
		assert.Nil(t, resp.Error)

		parentStored = aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored = aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Error)
		assert.Equal(t, child.Id(), childStored.Resource.Id())
		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
	})

	t.Run("Successfully connect nonexistent resources", func(t *testing.T) {
		parent := group
		child := user

		parentStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Resource)
		assert.ErrorIs(t, childStored.Error, acl.ErrNotFound)

		req := syncer.ConnectResourcesReq{
			ReqId:  "1",
			Parent: parent,
			Child:  child,
		}
		resp := aclStore.ConnectResources(acl.ConnectResourcesReq{
			Parent:   req.Parent,
			Child:    req.Child,
			Callback: outboxMessageCallback(req),
		})
		assert.Nil(t, resp.Error)

		parentStored = aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored = aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Error)
		assert.Equal(t, child.Id(), childStored.Resource.Id())
		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
	})

	t.Run("Successfully connect existing resources", func(t *testing.T) {
		parent := org
		child := user

		parentStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Error)
		assert.Equal(t, child.Id(), childStored.Resource.Id())
		assert.Equal(t, child.Kind(), childStored.Resource.Kind())

		req := syncer.ConnectResourcesReq{
			ReqId:  "1",
			Parent: parent,
			Child:  child,
		}
		resp := aclStore.ConnectResources(acl.ConnectResourcesReq{
			Parent:   req.Parent,
			Child:    req.Child,
			Callback: outboxMessageCallback(req),
		})
		assert.Nil(t, resp.Error)

		parentStored = aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored = aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Error)
		assert.Equal(t, child.Id(), childStored.Resource.Id())
		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
	})

	//disconnect user from the group
	t.Run("Disconnect resource (no orphan descendants)", func(t *testing.T) {
		parent := group
		child := user

		parentStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Error)
		assert.Equal(t, child.Id(), childStored.Resource.Id())
		assert.Equal(t, child.Kind(), childStored.Resource.Kind())

		req := syncer.DisconnectResourcesReq{
			ReqId:  "1",
			Parent: parent,
			Child:  child,
		}
		resp := aclStore.DisconnectResources(acl.DisconnectResourcesReq{
			Parent:   req.Parent,
			Child:    req.Child,
			Callback: outboxMessageCallback(req),
		})
		assert.Nil(t, resp.Error)

		orgStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   org.Id(),
			Kind: org.Kind(),
		})
		assert.Nil(t, orgStored.Error)
		assert.Equal(t, org.Id(), orgStored.Resource.Id())
		assert.Equal(t, org.Kind(), orgStored.Resource.Kind())
		userStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   user.Id(),
			Kind: user.Kind(),
		})
		assert.Nil(t, userStored.Error)
		assert.Equal(t, user.Id(), userStored.Resource.Id())
		assert.Equal(t, user.Kind(), userStored.Resource.Kind())
		groupStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   group.Id(),
			Kind: group.Kind(),
		})
		assert.Nil(t, groupStored.Error)
		assert.Equal(t, group.Id(), groupStored.Resource.Id())
		assert.Equal(t, group.Kind(), groupStored.Resource.Kind())
	})

	//disconnect user from the organization (user should be deleted)
	t.Run("Disconnect resource (orphan descendants)", func(t *testing.T) {
		parent := org
		child := user

		parentStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   parent.Id(),
			Kind: parent.Kind(),
		})
		assert.Nil(t, parentStored.Error)
		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
		childStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   child.Id(),
			Kind: child.Kind(),
		})
		assert.Nil(t, childStored.Error)
		assert.Equal(t, child.Id(), childStored.Resource.Id())
		assert.Equal(t, child.Kind(), childStored.Resource.Kind())

		req := syncer.DisconnectResourcesReq{
			ReqId:  "1",
			Parent: parent,
			Child:  child,
		}
		resp := aclStore.DisconnectResources(acl.DisconnectResourcesReq{
			Parent:   req.Parent,
			Child:    req.Child,
			Callback: outboxMessageCallback(req),
		})
		assert.Nil(t, resp.Error)

		orgStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   org.Id(),
			Kind: org.Kind(),
		})
		assert.Nil(t, orgStored.Error)
		assert.Equal(t, org.Id(), orgStored.Resource.Id())
		assert.Equal(t, org.Kind(), orgStored.Resource.Kind())
		userStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   user.Id(),
			Kind: user.Kind(),
		})
		assert.Nil(t, userStored.Resource)
		assert.ErrorIs(t, userStored.Error, acl.ErrNotFound)
		groupStored := aclStore.GetResource(acl.GetResourceReq{
			Id:   group.Id(),
			Kind: group.Kind(),
		})
		assert.Nil(t, groupStored.Error)
		assert.Equal(t, group.Id(), groupStored.Resource.Id())
		assert.Equal(t, group.Kind(), groupStored.Resource.Kind())
	})

	cleanUpAclStore(t)
}

// outboxMessageCallback is a helper function that generates an OutboxMessage based on a sync request
// outbox messages serve the purpose of atomically committing changes and publishing events afterwards
func outboxMessageCallback(req syncer.Request) func(error) *model.OutboxMessage {
	syncRespFactory := syncerpb.NewSyncRespOutboxMessage
	return func(err error) *model.OutboxMessage {
		if err != nil {
			return syncRespFactory(req.Id(), err.Error(), false)
		}
		return syncRespFactory(req.Id(), "", true)
	}
}

func setUpAclStore(t *testing.T) {
	if aclStore != nil {
		return
	}

	c, err := setupNeo4jContainer(context.Background())
	if err != nil {
		t.Error(err)
	}
	txManager, err = neo4jstore.NewTransactionManager(c.uri, c.dbName)
	if err != nil {
		t.Error(err)
	}

	aclStore = neo4jstore.NewAclStore(txManager)
}

func cleanUpAclStore(t *testing.T) {
	if txManager == nil {
		return
	}

	_, err := txManager.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				"MATCH (n) DETACH DELETE n",
				map[string]interface{}{},
			)
			if err != nil {
				return nil, err
			}
			return nil, result.Err()
		})

	if err != nil {
		t.Error(err)
	}
}
