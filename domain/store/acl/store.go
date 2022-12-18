package acl

type Store interface {
	CreateResource(req CreateResourceReq) SyncResp
	DeleteResource(req DeleteResourceReq) SyncResp
	GetResource(req GetResourceReq) GetResourceResp
	CreateAttribute(req CreateAttributeReq) SyncResp
	UpdateAttribute(req UpdateAttributeReq) SyncResp
	DeleteAttribute(req DeleteAttributeReq) SyncResp
	CreateAggregationRel(req CreateAggregationRelReq) SyncResp
	DeleteAggregationRel(req DeleteAggregationRelReq) SyncResp
	CreateCompositionRel(req CreateCompositionRelReq) SyncResp
	DeleteCompositionRel(req DeleteCompositionRelReq) SyncResp
	CreatePermission(req CreatePermissionReq) SyncResp
	DeletePermission(req DeletePermissionReq) SyncResp
	GetPermissionHierarchy(req GetPermissionHierarchyReq) GetPermissionHierarchyResp
}
