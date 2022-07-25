package neo4j

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"sort"
)

func getAttributes(cypherResults interface{}) []model.Attribute {
	attributes := make([]model.Attribute, 0)
	fmt.Println(cypherResults)
	for _, result := range cypherResults.([]interface{}) {
		fmt.Println(result)
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

func getPermission(cypherResult interface{}) model.Permission {
	permMap := cypherResult.(map[string]interface{})
	return model.NewPermission(
		permMap["name"].(string),
		model.PermissionKind(permMap["kind"].(int64)),
		model.NewCondition(
			permMap["condition"].(string)))
}

func sortByDistanceAsc(m map[int]model.PermissionList) model.PermissionHierarchy {
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
