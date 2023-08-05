package acl

type Store interface {
	CreateResource(req CreateResourceReq) SyncResp
	DeleteResource(req DeleteResourceReq) SyncResp
	GetResource(req GetResourceReq) GetResourceResp
	PutAttribute(req PutAttributeReq) SyncResp
	DeleteAttribute(req DeleteAttributeReq) SyncResp
	CreateInheritanceRel(req CreateInheritanceRelReq) SyncResp
	DeleteInheritanceRel(req DeleteInheritanceRelReq) SyncResp
	CreatePolicy(req CreatePolicyReq) SyncResp
	DeletePolicy(req DeletePolicyReq) SyncResp
	GetPermissionHierarchy(req GetPermissionHierarchyReq) GetPermissionHierarchyResp
}
