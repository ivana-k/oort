package domain

type AdministrationReqMarshaller interface {
	Marshal(msg AdministrationReq) ([]byte, error)
	Unmarshal(bytes []byte) (*AdministrationReq, error)
}

//type AttributesMarshaller interface {
//	Marshal(attributes []Attribute) ([]byte, error)
//	Unmarshal(bytes []byte) ([]Attribute, error)
//}
//
//type AuthorizationRespMarshaller interface {
//	Marshal(resp AuthorizationResp) ([]byte, error)
//	Unmarshal(bytes []byte) (*AuthorizationResp, error)
//}
