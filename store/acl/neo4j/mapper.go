package neo4j

import (
	"errors"
	"github.com/c12s/oort/domain/model"
)

func getResource(cypherResult interface{}) *model.Resource {
	resource := model.NewResource("", "")
	resource.Attributes = make([]model.Attribute, 0)
	attrs := cypherResult.([]interface{})[1].([]interface{})
	for _, attr := range attrs {
		a := attr.(map[string]interface{})
		name := a["name"].(string)
		kind := model.AttributeKind(a["kind"].(int64))
		value := a["value"]
		if name == "id" {
			resource.SetId(value.(string))
		}
		if name == "kind" {
			resource.SetKind(value.(string))
		}
		attribute := model.NewAttribute(model.NewAttributeId(name), kind, value)
		resource.Attributes = append(resource.Attributes, attribute)
	}
	return &resource
}

func getHierarchy(cypherResult interface{}) (model.PermissionHierarchy, error) {
	recordList, ok := cypherResult.([]interface{})
	if !ok {
		return model.PermissionHierarchy{}, errors.New("invalid resp format")
	}

	hierarchy := make(map[model.PermissionPriority]model.PermissionObjHierarchy)
	for _, record := range recordList {
		recordElems, ok := record.([]interface{})
		if !ok {
			return model.PermissionHierarchy{}, errors.New("invalid resp format")
		}

		permName, ok := recordElems[0].(string)
		if !ok {
			return model.PermissionHierarchy{}, errors.New("invalid record elem type - perm name")
		}
		permKindInt, ok := recordElems[1].(int64)
		if !ok {
			return model.PermissionHierarchy{}, errors.New("invalid record elem type - perm kind")
		}
		permKind := model.PermissionKind(permKindInt)
		permCond, ok := recordElems[2].(string)
		if !ok {
			return model.PermissionHierarchy{}, errors.New("invalid record elem type - perm cond")
		}
		subPriorityInt, ok := recordElems[3].(int64)
		if !ok {
			return model.PermissionHierarchy{}, errors.New("invalid record elem type - perm sub priority")
		}
		subPriority := model.PermissionPriority(subPriorityInt)
		objPriorityInt, ok := recordElems[4].(int64)
		if !ok {
			return model.PermissionHierarchy{}, errors.New("invalid record elem type - perm obj priority")
		}
		objPriority := model.PermissionPriority(objPriorityInt)

		// kreiraj dozvolu
		cond, err := model.NewCondition(permCond)
		if err != nil {
			return model.PermissionHierarchy{}, errors.New("invalid condition")
		}
		perm := model.NewPermission(permName, permKind, *cond)
		// proveri kom obj hierarchy elem pripada, ako ga nema kreiraj
		_, ok = hierarchy[subPriority]
		if !ok {
			hierarchy[subPriority] = make(map[model.PermissionPriority]model.PermissionLevel)
		}
		objHierarchy := hierarchy[subPriority]
		// proveri kom perm level-u (unutar obj hierarchy) elem pripada, ako ga nema kreiraj
		_, ok = objHierarchy[objPriority]
		if !ok {
			objHierarchy[objPriority] = make([]model.Permission, 0)
		}
		// perm level-u dodaj perm
		objHierarchy[objPriority] = append(objHierarchy[objPriority], perm)
		// izmeni hierarchy, dodeli mu novi obj hierarchy
		hierarchy[subPriority] = objHierarchy
	}
	return hierarchy, nil
}
