package domain

type RHABACRepo interface {
	CreateResource(req CreateResourceReq) AdministrationResp
	DeleteResource(req DeleteResourceReq) AdministrationResp
	GetResource(req GetResourceReq) GetResourceResp
	PutAttribute(req PutAttributeReq) AdministrationResp
	DeleteAttribute(req DeleteAttributeReq) AdministrationResp
	CreateInheritanceRel(req CreateInheritanceRelReq) AdministrationResp
	DeleteInheritanceRel(req DeleteInheritanceRelReq) AdministrationResp
	CreatePolicy(req CreatePolicyReq) AdministrationResp
	DeletePolicy(req DeletePolicyReq) AdministrationResp
	GetPermissionHierarchy(req GetPermissionHierarchyReq) GetPermissionHierarchyResp
}

type CreateResourceReq struct {
	Resource Resource
}

type DeleteResourceReq struct {
	Resource Resource
}

type GetResourceReq struct {
	Resource Resource
}

type PutAttributeReq struct {
	Resource  Resource
	Attribute Attribute
}

type DeleteAttributeReq struct {
	Resource    Resource
	AttributeId AttributeId
}

type GetAttributeReq struct {
	Resource Resource
}

type CreateInheritanceRelReq struct {
	From Resource
	To   Resource
}

type DeleteInheritanceRelReq struct {
	From Resource
	To   Resource
}

type CreatePolicyReq struct {
	SubjectScope,
	ObjectScope Resource
	Permission Permission
}

type DeletePolicyReq struct {
	SubjectScope,
	ObjectScope Resource
	Permission Permission
}

type GetPermissionHierarchyReq struct {
	Subject,
	Object Resource
	PermissionName string
}

type AdministrationResp struct {
	Error error
}

type GetAttributeResp struct {
	Attributes []Attribute
	Error      error
}

type GetResourceResp struct {
	Resource *Resource
	Error    error
}

type GetPermissionHierarchyResp struct {
	Hierarchy PermissionHierarchy
	Error     error
}

type AuthorizationReq struct {
	Subject,
	Object Resource
	PermissionName string
	Env            []Attribute
}

type AuthorizationResp struct {
	Authorized bool
	Error      error
}
