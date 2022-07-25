package checkerpb

import (
	"github.com/c12s/oort/domain/model/checker"
)

func (x *CheckPermissionReq) MapToDomain() checker.CheckPermissionReq {
	return checker.CheckPermissionReq{
		Principal:      x.Principal.MapToDomain(),
		Resource:       x.Resource.MapToDomain(),
		PermissionName: x.PermissionName,
	}
}
