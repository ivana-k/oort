package syncerpb

import (
	"github.com/c12s/oort/domain/model/syncer"
)

func (x *ConnectResourcesReq) MapToDomain() syncer.ConnectResourcesReq {
	return syncer.ConnectResourcesReq{
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}
}

func (x *DisconnectResourcesReq) MapToDomain() syncer.DisconnectResourcesReq {
	return syncer.DisconnectResourcesReq{
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}
}

func (x *UpsertAttributeReq) MapToDomain() (syncer.UpsertAttributeReq, error) {
	attr, err := x.Attribute.MapToDomain()
	if err != nil {
		return syncer.UpsertAttributeReq{}, err
	}
	return syncer.UpsertAttributeReq{
		Resource:  x.Resource.MapToDomain(),
		Attribute: attr,
	}, nil
}

func (x *RemoveAttributeReq) MapToDomain() syncer.RemoveAttributeReq {
	return syncer.RemoveAttributeReq{
		Resource:    x.Resource.MapToDomain(),
		AttributeId: x.AttributeId.MapToDomain(),
	}
}

func (x *InsertPermissionReq) MapToDomain() syncer.InsertPermissionReq {
	return syncer.InsertPermissionReq{
		Principal:  x.Principal.MapToDomain(),
		Resource:   x.Resource.MapToDomain(),
		Permission: x.Permission.MapToDomain(),
	}
}

func (x *RemovePermissionReq) MapToDomain() syncer.RemovePermissionReq {
	return syncer.RemovePermissionReq{
		Principal:  x.Principal.MapToDomain(),
		Resource:   x.Resource.MapToDomain(),
		Permission: x.Permission.MapToDomain(),
	}
}
