package domain

type RHABACStore interface {
	CreateResource(req CreateResourceReq, generator OutboxMsgGenerator) AdministrationResp
	DeleteResource(req DeleteResourceReq, generator OutboxMsgGenerator) AdministrationResp
	GetResource(req GetResourceReq) GetResourceResp
	PutAttribute(req PutAttributeReq, generator OutboxMsgGenerator) AdministrationResp
	DeleteAttribute(req DeleteAttributeReq, generator OutboxMsgGenerator) AdministrationResp
	CreateInheritanceRel(req CreateInheritanceRelReq, generator OutboxMsgGenerator) AdministrationResp
	DeleteInheritanceRel(req DeleteInheritanceRelReq, generator OutboxMsgGenerator) AdministrationResp
	CreatePolicy(req CreatePolicyReq, generator OutboxMsgGenerator) AdministrationResp
	DeletePolicy(req DeletePolicyReq, generator OutboxMsgGenerator) AdministrationResp
	GetPermissionHierarchy(req GetPermissionHierarchyReq) GetPermissionHierarchyResp
}

type CreateResourceReq struct {
	Id       string
	Resource Resource
}

type DeleteResourceReq struct {
	Id       string
	Resource Resource
}

type GetResourceReq struct {
	Resource Resource
}

type PutAttributeReq struct {
	Id        string
	Resource  Resource
	Attribute Attribute
}

type DeleteAttributeReq struct {
	Id          string
	Resource    Resource
	AttributeId AttributeId
}

type GetAttributeReq struct {
	Resource Resource
}

type CreateInheritanceRelReq struct {
	Id   string
	From Resource
	To   Resource
}

type DeleteInheritanceRelReq struct {
	Id   string
	From Resource
	To   Resource
}

type CreatePolicyReq struct {
	Id string
	SubjectScope,
	ObjectScope Resource
	Permission Permission
}

type DeletePolicyReq struct {
	Id string
	SubjectScope,
	ObjectScope Resource
	Permission Permission
}

type GetPermissionHierarchyReq struct {
	Subject,
	Object Resource
	PermissionName string
}

type AdministrationReqKind int

const (
	CreateResource AdministrationReqKind = iota
	DeleteResource
	PutAttribute
	DeleteAttribute
	CreateInheritanceRel
	DeleteInheritanceRel
	CreatePolicy
	DeletePolicy
)

type AdministrationReq struct {
	ReqKind AdministrationReqKind
	Request interface{}
}

type OutboxMsgGenerator func(error) *OutboxMessage

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
	Allowed bool
	Error   error
}
