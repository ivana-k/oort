package test

//import (
//	"github.com/c12s/oort/cmd/domain/repos/acl"
//	neo4jstore "github.com/c12s/oort/cmd/repos/rhabac/neo4j"
//)
//
//var (
//	AclStore  acl.Store
//	TxManager *neo4jstore.TransactionManager
//)

//
//func TestResourceConnections(t *testing.T) {
//	setUpAclStore(t)
//
//	org := model.NewResource("org1", "org")
//	group := model.NewResource("g1", "group")
//	user := model.NewResource("u1", "user")
//
//	t.Run("Successfully connect nonexistent resources", func(t *testing.T) {
//		parent := org
//		child := group
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Resource)
//		assert.ErrorIs(t, parentStored.Error, rhabac.ErrNotFound)
//		childStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Resource)
//		assert.ErrorIs(t, childStored.Error, rhabac.ErrNotFound)
//
//		req := syncer.ConnectResourcesReq{
//			ReqId:  "1",
//			Parent: parent,
//			Child:  child,
//		}
//		resp := aclStore.ConnectResources(rhabac.ConnectResourcesReq{
//			Parent:   req.Parent,
//			Child:    req.Child,
//			Callback: outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		parentStored = aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored = aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Error)
//		assert.Equal(t, child.Id(), childStored.Resource.Id())
//		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
//	})
//
//	t.Run("Successfully connect nonexistent resources", func(t *testing.T) {
//		parent := group
//		child := user
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Resource)
//		assert.ErrorIs(t, childStored.Error, rhabac.ErrNotFound)
//
//		req := syncer.ConnectResourcesReq{
//			ReqId:  "1",
//			Parent: parent,
//			Child:  child,
//		}
//		resp := aclStore.ConnectResources(rhabac.ConnectResourcesReq{
//			Parent:   req.Parent,
//			Child:    req.Child,
//			Callback: outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		parentStored = aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored = aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Error)
//		assert.Equal(t, child.Id(), childStored.Resource.Id())
//		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
//	})
//
//	t.Run("Successfully connect existing resources", func(t *testing.T) {
//		parent := org
//		child := user
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Error)
//		assert.Equal(t, child.Id(), childStored.Resource.Id())
//		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
//
//		req := syncer.ConnectResourcesReq{
//			ReqId:  "1",
//			Parent: parent,
//			Child:  child,
//		}
//		resp := aclStore.ConnectResources(rhabac.ConnectResourcesReq{
//			Parent:   req.Parent,
//			Child:    req.Child,
//			Callback: outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		parentStored = aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored = aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Error)
//		assert.Equal(t, child.Id(), childStored.Resource.Id())
//		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
//	})
//
//	//disconnect user from the group
//	t.Run("Disconnect resource (no orphan descendants)", func(t *testing.T) {
//		parent := group
//		child := user
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Error)
//		assert.Equal(t, child.Id(), childStored.Resource.Id())
//		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
//
//		req := syncer.DisconnectResourcesReq{
//			ReqId:  "1",
//			Parent: parent,
//			Child:  child,
//		}
//		resp := aclStore.DisconnectResources(rhabac.DisconnectResourcesReq{
//			Parent:   req.Parent,
//			Child:    req.Child,
//			Callback: outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		orgStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   org.Id(),
//			Kind: org.Kind(),
//		})
//		assert.Nil(t, orgStored.Error)
//		assert.Equal(t, org.Id(), orgStored.Resource.Id())
//		assert.Equal(t, org.Kind(), orgStored.Resource.Kind())
//		userStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   user.Id(),
//			Kind: user.Kind(),
//		})
//		assert.Nil(t, userStored.Error)
//		assert.Equal(t, user.Id(), userStored.Resource.Id())
//		assert.Equal(t, user.Kind(), userStored.Resource.Kind())
//		groupStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   group.Id(),
//			Kind: group.Kind(),
//		})
//		assert.Nil(t, groupStored.Error)
//		assert.Equal(t, group.Id(), groupStored.Resource.Id())
//		assert.Equal(t, group.Kind(), groupStored.Resource.Kind())
//	})
//
//	//disconnect user from the organization (user should be deleted)
//	t.Run("Disconnect resource (orphan descendants)", func(t *testing.T) {
//		parent := org
//		child := user
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Error)
//		assert.Equal(t, parent.Id(), parentStored.Resource.Id())
//		assert.Equal(t, parent.Kind(), parentStored.Resource.Kind())
//		childStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Error)
//		assert.Equal(t, child.Id(), childStored.Resource.Id())
//		assert.Equal(t, child.Kind(), childStored.Resource.Kind())
//
//		req := syncer.DisconnectResourcesReq{
//			ReqId:  "1",
//			Parent: parent,
//			Child:  child,
//		}
//		resp := aclStore.DisconnectResources(rhabac.DisconnectResourcesReq{
//			Parent:   req.Parent,
//			Child:    req.Child,
//			Callback: outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		orgStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   org.Id(),
//			Kind: org.Kind(),
//		})
//		assert.Nil(t, orgStored.Error)
//		assert.Equal(t, org.Id(), orgStored.Resource.Id())
//		assert.Equal(t, org.Kind(), orgStored.Resource.Kind())
//		userStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   user.Id(),
//			Kind: user.Kind(),
//		})
//		assert.Nil(t, userStored.Resource)
//		assert.ErrorIs(t, userStored.Error, rhabac.ErrNotFound)
//		groupStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   group.Id(),
//			Kind: group.Kind(),
//		})
//		assert.Nil(t, groupStored.Error)
//		assert.Equal(t, group.Id(), groupStored.Resource.Id())
//		assert.Equal(t, group.Kind(), groupStored.Resource.Kind())
//	})
//
//	cleanUpAclStore(t)
//}
//
//func TestResourceAttributes(t *testing.T) {
//	setUpAclStore(t)
//
//	org := model.NewResource("org1", "org")
//	user := model.NewResource("u1", "user")
//	group := model.NewResource("g1", "group")
//	username := model.NewAttribute(model.NewAttributeId("name", model.String), "pera")
//	username2 := model.NewAttribute(model.NewAttributeId("name", model.String), "mika")
//	username3 := model.NewAttribute(model.NewAttributeId("name", model.Int64), int64(123))
//
//	t.Run("successfully insert attribute of an existing resource", func(t *testing.T) {
//		parent := org
//		child := user
//		attribute := username
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   parent.Id(),
//			Kind: parent.Kind(),
//		})
//		assert.Nil(t, parentStored.Resource)
//		assert.ErrorIs(t, parentStored.Error, rhabac.ErrNotFound)
//		childStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   child.Id(),
//			Kind: child.Kind(),
//		})
//		assert.Nil(t, childStored.Resource)
//		assert.ErrorIs(t, childStored.Error, rhabac.ErrNotFound)
//
//		req := syncer.ConnectResourcesReq{
//			ReqId:  "1",
//			Parent: parent,
//			Child:  child,
//		}
//		resp := aclStore.ConnectResources(rhabac.ConnectResourcesReq{
//			Parent:   req.Parent,
//			Child:    req.Child,
//			Callback: outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		upsertReq := syncer.UpsertAttributeReq{
//			ReqId:     "1",
//			Resource:  child,
//			Attribute: attribute,
//		}
//		upsertResp := aclStore.UpsertAttribute(rhabac.UpsertAttributeReq{
//			Resource:  upsertReq.Resource,
//			Attribute: upsertReq.Attribute,
//			Callback:  outboxMessageCallback(upsertReq),
//		})
//		assert.Nil(t, upsertResp.Error)
//
//		attrs := aclStore.GetAttributes(rhabac.GetAttributeReq{
//			Resource: upsertReq.Resource,
//		})
//		assert.Nil(t, attrs.Error)
//		assert.True(t, containsAttribute(attrs.Attributes, attribute))
//	})
//
//	t.Run("unsuccessful insertion attempt, resource does not exist", func(t *testing.T) {
//		resource := group
//		attribute := username
//
//		parentStored := aclStore.GetResource(rhabac.GetResourceReq{
//			Id:   resource.Id(),
//			Kind: resource.Kind(),
//		})
//		assert.Nil(t, parentStored.Resource)
//		assert.ErrorIs(t, parentStored.Error, rhabac.ErrNotFound)
//
//		upsertReq := syncer.UpsertAttributeReq{
//			ReqId:     "2",
//			Resource:  resource,
//			Attribute: attribute,
//		}
//		upsertResp := aclStore.UpsertAttribute(rhabac.UpsertAttributeReq{
//			Resource:  upsertReq.Resource,
//			Attribute: upsertReq.Attribute,
//			Callback:  outboxMessageCallback(upsertReq),
//		})
//		assert.Nil(t, upsertResp.Error)
//
//		attrs := aclStore.GetAttributes(rhabac.GetAttributeReq{
//			Resource: upsertReq.Resource,
//		})
//		for _, attr := range attrs.Attributes {
//			assert.NotEqual(t, attr, attribute)
//		}
//	})
//
//	t.Run("successfully update an existing attribute (type unchanged)", func(t *testing.T) {
//		resource := user
//		attribute := username
//		updatedAttribute := username2
//
//		oldAttrResp := aclStore.GetAttributes(rhabac.GetAttributeReq{
//			Resource: resource,
//		})
//		assert.Nil(t, oldAttrResp.Error)
//		assert.True(t, containsAttribute(oldAttrResp.Attributes, attribute))
//
//		upsertReq := syncer.UpsertAttributeReq{
//			ReqId:     "2",
//			Resource:  resource,
//			Attribute: updatedAttribute,
//		}
//		upsertResp := aclStore.UpsertAttribute(rhabac.UpsertAttributeReq{
//			Resource:  upsertReq.Resource,
//			Attribute: upsertReq.Attribute,
//			Callback:  outboxMessageCallback(upsertReq),
//		})
//		assert.Nil(t, upsertResp.Error)
//
//		attrs := aclStore.GetAttributes(rhabac.GetAttributeReq{
//			Resource: upsertReq.Resource,
//		})
//		assert.True(t, containsAttribute(attrs.Attributes, updatedAttribute))
//		assert.False(t, containsAttribute(attrs.Attributes, attribute))
//	})
//
//	t.Run("successfully update an existing attribute (type changed)", func(t *testing.T) {
//		resource := user
//		attribute := username2
//		updatedAttribute := username3
//
//		oldAttrResp := aclStore.GetAttributes(rhabac.GetAttributeReq{
//			Resource: resource,
//		})
//		assert.Nil(t, oldAttrResp.Error)
//		assert.True(t, containsAttribute(oldAttrResp.Attributes, attribute))
//
//		upsertReq := syncer.UpsertAttributeReq{
//			ReqId:     "3",
//			Resource:  resource,
//			Attribute: updatedAttribute,
//		}
//		upsertResp := aclStore.UpsertAttribute(rhabac.UpsertAttributeReq{
//			Resource:  upsertReq.Resource,
//			Attribute: upsertReq.Attribute,
//			Callback:  outboxMessageCallback(upsertReq),
//		})
//		assert.Nil(t, upsertResp.Error)
//
//		attrs := aclStore.GetAttributes(rhabac.GetAttributeReq{
//			Resource: upsertReq.Resource,
//		})
//		assert.True(t, containsAttribute(attrs.Attributes, updatedAttribute))
//		assert.False(t, containsAttribute(attrs.Attributes, attribute))
//	})
//
//	cleanUpAclStore(t)
//}
//
//func TestPermissions(t *testing.T) {
//	setUpAclStore(t)
//
//	org := model.NewResource("org1", "org")
//	user := model.NewResource("u1", "user")
//	region := model.NewResource("r1", "region")
//	cluster := model.NewResource("c1", "cluster")
//	permission := model.NewPermission("cluster.get", model.PermissionKindAllow, model.Condition{})
//	permission2 := model.NewPermission("cluster.create", model.PermissionKindDeny, model.Condition{})
//	permission3 := model.NewPermission("cluster.delete", model.PermissionKindDeny, model.Condition{})
//
//	req := syncer.ConnectResourcesReq{
//		ReqId:  "1",
//		Parent: org,
//		Child:  user,
//	}
//	resp := aclStore.ConnectResources(rhabac.ConnectResourcesReq{
//		Parent:   req.Parent,
//		Child:    req.Child,
//		Callback: outboxMessageCallback(req),
//	})
//	assert.Nil(t, resp.Error)
//	req = syncer.ConnectResourcesReq{
//		ReqId:  "1",
//		Parent: region,
//		Child:  cluster,
//	}
//	resp = aclStore.ConnectResources(rhabac.ConnectResourcesReq{
//		Parent:   req.Parent,
//		Child:    req.Child,
//		Callback: outboxMessageCallback(req),
//	})
//	assert.Nil(t, resp.Error)
//	resourceResp := aclStore.GetResource(rhabac.GetResourceReq{
//		Id:   org.Id(),
//		Kind: org.Kind(),
//	})
//	assert.Equal(t, *resourceResp.Resource, org)
//	resourceResp = aclStore.GetResource(rhabac.GetResourceReq{
//		Id:   user.Id(),
//		Kind: user.Kind(),
//	})
//	assert.Equal(t, *resourceResp.Resource, user)
//	resourceResp = aclStore.GetResource(rhabac.GetResourceReq{
//		Id:   region.Id(),
//		Kind: region.Kind(),
//	})
//	assert.Equal(t, *resourceResp.Resource, region)
//	resourceResp = aclStore.GetResource(rhabac.GetResourceReq{
//		Id:   cluster.Id(),
//		Kind: cluster.Kind(),
//	})
//	assert.Equal(t, *resourceResp.Resource, cluster)
//
//	t.Run("direct permission assignment, no inheritance", func(t *testing.T) {
//		principal := user
//		resource := cluster
//		permission := permission
//
//		req := syncer.InsertPermissionReq{
//			ReqId:      "1",
//			Principal:  principal,
//			Resource:   resource,
//			Permission: permission,
//		}
//		resp := aclStore.InsertPermission(rhabac.InsertPermissionReq{
//			Principal:  req.Principal,
//			Resource:   req.Resource,
//			Permission: req.Permission,
//			Callback:   outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		permissionsResp := aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      req.Principal,
//			Resource:       req.Resource,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.True(t, containsPermission(permissionsResp.Hierarchy, permission))
//		//newly created permission didn't affect any other principal-resource combinations
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      user,
//			Resource:       region,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      org,
//			Resource:       region,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      org,
//			Resource:       cluster,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//	})
//	t.Run("subject-side inherited permission", func(t *testing.T) {
//		principal := org
//		resource := cluster
//		permission := permission2
//
//		req := syncer.InsertPermissionReq{
//			ReqId:      "2",
//			Principal:  principal,
//			Resource:   resource,
//			Permission: permission,
//		}
//		resp := aclStore.InsertPermission(rhabac.InsertPermissionReq{
//			Principal:  req.Principal,
//			Resource:   req.Resource,
//			Permission: req.Permission,
//			Callback:   outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		permissionsResp := aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      req.Principal,
//			Resource:       req.Resource,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.True(t, containsPermission(permissionsResp.Hierarchy, permission))
//		// user inherits permissions from org
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      user,
//			Resource:       cluster,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.True(t, containsPermission(permissionsResp.Hierarchy, permission))
//		// no inheritance
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      org,
//			Resource:       region,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      user,
//			Resource:       region,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//	})
//	t.Run("object-side inherited permission", func(t *testing.T) {
//		principal := user
//		resource := region
//		permission := permission3
//
//		req := syncer.InsertPermissionReq{
//			ReqId:      "2",
//			Principal:  principal,
//			Resource:   resource,
//			Permission: permission,
//		}
//		resp := aclStore.InsertPermission(rhabac.InsertPermissionReq{
//			Principal:  req.Principal,
//			Resource:   req.Resource,
//			Permission: req.Permission,
//			Callback:   outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		permissionsResp := aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      req.Principal,
//			Resource:       req.Resource,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.True(t, containsPermission(permissionsResp.Hierarchy, permission))
//		// region-scoped permission -> cluster-scoped permission
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      user,
//			Resource:       cluster,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.True(t, containsPermission(permissionsResp.Hierarchy, permission))
//		// no inheritance
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      org,
//			Resource:       region,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//		permissionsResp = aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      org,
//			Resource:       cluster,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//	})
//	t.Run("successfully deleting permissions", func(t *testing.T) {
//		principal := user
//		resource := region
//		permission := permission3
//
//		req := syncer.RemovePermissionReq{
//			ReqId:      "2",
//			Principal:  principal,
//			Resource:   resource,
//			Permission: permission,
//		}
//		resp := aclStore.RemovePermission(rhabac.RemovePermissionReq{
//			Principal:  req.Principal,
//			Resource:   req.Resource,
//			Permission: req.Permission,
//			Callback:   outboxMessageCallback(req),
//		})
//		assert.Nil(t, resp.Error)
//
//		permissionsResp := aclStore.GetPermissionByPrecedence(rhabac.GetPermissionReq{
//			Principal:      req.Principal,
//			Resource:       req.Resource,
//			PermissionName: req.Permission.Name(),
//		})
//		assert.Nil(t, permissionsResp.Error)
//		assert.False(t, containsPermission(permissionsResp.Hierarchy, permission))
//	})
//
//	cleanUpAclStore(t)
//}
//
//// outboxMessageCallback is a helper function that generates an OutboxMessage based on a sync request
//// outbox messages serve the purpose of atomically committing changes and publishing events afterwards
//func outboxMessageCallback(req syncer.Request) func(error) *model.OutboxMessage {
//	syncRespFactory := syncerpb.NewSyncRespOutboxMessage
//	return func(err error) *model.OutboxMessage {
//		if err != nil {
//			return syncRespFactory(req.Id(), err.Error(), false)
//		}
//		return syncRespFactory(req.Id(), "", true)
//	}
//}
//
//func containsAttribute(list []model.Attribute, attribute model.Attribute) bool {
//	for _, attr := range list {
//		if attr == attribute {
//			return true
//		}
//	}
//	return false
//}
//
//func containsPermission(hierarchy model.PermissionHierarchy, permission model.Permission) bool {
//	for _, level := range hierarchy {
//		for _, perm := range level {
//			if perm == permission {
//				return true
//			}
//		}
//	}
//	return false
//}
