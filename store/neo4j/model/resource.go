package model

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
)

type Resource struct {
	Resource model.Resource
}

func (r Resource) Properties(resourceVar string) (cypher string, params map[string]interface{}) {
	params = map[string]interface{}{}
	for key, value := range r.Resource.GetArgs() {
		cypher += fmt.Sprintf("%s.%s = $%s, ", resourceVar, key, key)
		params[key] = value
	}
	cypher = cypher[:len(cypher)-2]
	return
}

func (r Resource) ResourcePattern(resourceVar string) string {
	return fmt.Sprintf("(%s:%s)", resourceVar, resourceLabel)
}

func (r Resource) IdentityPattern(resourceVar string) string {
	return fmt.Sprintf("(%s:%s:%s)", resourceVar, resourceLabel, identityLabel)
}
