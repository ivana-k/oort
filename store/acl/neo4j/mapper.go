package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"sort"
)

func getAttributes(cypherResults interface{}) []model.Attribute {
	attributes := make([]model.Attribute, 0)
	for _, result := range cypherResults.([]interface{}) {
		attrMap := result.(map[string]interface{})
		attr := model.NewAttribute(
			model.NewAttributeId(
				attrMap["name"].(string),
				model.AttributeKind(attrMap["kind"].(int64))),
			attrMap["value"].([]byte))
		attributes = append(attributes, attr)
	}
	return attributes
}

func getResource(cypherResult interface{}) *model.Resource {
	attrMap := cypherResult.(map[string]interface{})
	id := attrMap["id"].(string)
	kind := attrMap["kind"].(string)
	resource := model.NewResource(id, kind)
	return &resource
}

func getPermission(cypherResult interface{}) (model.Permission, error) {
	permMap := cypherResult.(map[string]interface{})
	condition, err := model.NewCondition(permMap["condition"].(string))
	if err != nil {
		return model.Permission{}, err
	}
	return model.NewPermission(
		permMap["name"].(string),
		model.PermissionKind(permMap["kind"].(int64)),
		*condition), nil
}

func sortByDistanceAsc(m map[int]model.PermissionLevel) model.PermissionHierarchy {
	keys := make([]int, 0)
	result := make(model.PermissionHierarchy, 0)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for key := range keys {
		result = append(result, m[key])
	}
	return result
}
