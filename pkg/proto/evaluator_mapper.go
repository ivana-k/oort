package proto

import (
	"github.com/c12s/oort/internal/domain"
	"log"
)

func (x *AuthorizationReq) ToDomain() (*domain.AuthorizationReq, error) {
	envAttributes := make([]domain.Attribute, len(x.EnvAttributes))
	for i, attr := range x.EnvAttributes {
		domainAttr, err := attr.ToDomain()
		if err != nil {
			log.Println(err)
			continue
		}
		envAttributes[i] = *domainAttr
	}
	sub, err := x.Subject.ToDomain()
	if err != nil {
		return nil, err
	}
	obj, err := x.Object.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.AuthorizationReq{
		Subject:        *sub,
		Object:         *obj,
		PermissionName: x.PermissionName,
		Env:            envAttributes,
	}, nil
}
