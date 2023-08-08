package neo4j

import (
	"errors"
	"fmt"
	"github.com/c12s/oort/internal/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

func getResource(cypherResult interface{}) *domain.Resource {
	resource, err := domain.NewResource("", "")
	if err != nil {
		return nil
	}
	resource.Attributes = make([]domain.Attribute, 0)
	attrs := cypherResult.([]*neo4j.Record)[0].Values[1].([]interface{})
	for _, attr := range attrs {
		a := attr.(map[string]interface{})
		name := a["name"].(string)
		kind := domain.AttributeKind(a["kind"].(int64))
		value := a["value"]
		if name == "id" {
			resource.SetId(value.(string))
		}
		if name == "kind" {
			resource.SetKind(value.(string))
		}
		attrId, err := domain.NewAttributeId(name)
		if err != nil {
			return nil
		}
		attribute, err := domain.NewAttribute(*attrId, kind, value)
		if err != nil {
			return nil
		}
		resource.Attributes = append(resource.Attributes, *attribute)
	}
	return resource
}

func getHierarchy(cypherResult interface{}) (domain.PermissionHierarchy, error) {
	records, ok := cypherResult.([]*neo4j.Record)
	log.Println(len(records))
	if !ok {
		return domain.PermissionHierarchy{}, errors.New("invalid resp format")
	}

	hierarchy := make(map[domain.PermissionPriority]domain.PermissionObjHierarchy)
	for _, record := range records {
		recordElems := record.Values
		if !ok {
			return domain.PermissionHierarchy{}, errors.New("invalid resp format")
		}

		permName, ok := recordElems[0].(string)
		if !ok {
			return domain.PermissionHierarchy{}, errors.New("invalid record elem type - perm name")
		}
		permKindInt, ok := recordElems[1].(int64)
		if !ok {
			return domain.PermissionHierarchy{}, errors.New("invalid record elem type - perm kind")
		}
		permKind := domain.PermissionKind(permKindInt)
		permCond, ok := recordElems[2].(string)
		if !ok {
			return domain.PermissionHierarchy{}, errors.New("invalid record elem type - perm cond")
		}
		subPriorityInt, ok := recordElems[3].(int64)
		if !ok {
			log.Println(fmt.Sprintf("%T", recordElems[3]))
			return domain.PermissionHierarchy{}, errors.New("invalid record elem type - perm sub priority")
		}
		subPriority := domain.PermissionPriority(subPriorityInt)
		objPriorityInt, ok := recordElems[4].(int64)
		if !ok {
			return domain.PermissionHierarchy{}, errors.New("invalid record elem type - perm obj priority")
		}
		objPriority := domain.PermissionPriority(objPriorityInt)

		// kreiraj dozvolu
		cond, err := domain.NewCondition(permCond)
		if err != nil {
			return domain.PermissionHierarchy{}, errors.New("invalid condition")
		}
		perm, err := domain.NewPermission(permName, permKind, *cond)
		if err != nil {
			return nil, err
		}
		// proveri kom obj hierarchy elem pripada, ako ga nema kreiraj
		_, ok = hierarchy[subPriority]
		if !ok {
			hierarchy[subPriority] = make(map[domain.PermissionPriority]domain.PermissionLevel)
		}
		objHierarchy := hierarchy[subPriority]
		// proveri kom perm level-u (unutar obj hierarchy) elem pripada, ako ga nema kreiraj
		_, ok = objHierarchy[objPriority]
		if !ok {
			objHierarchy[objPriority] = make([]domain.Permission, 0)
		}
		// perm level-u dodaj perm
		objHierarchy[objPriority] = append(objHierarchy[objPriority], *perm)
		// izmeni hierarchy, dodeli mu novi obj hierarchy
		hierarchy[subPriority] = objHierarchy
	}
	return hierarchy, nil
}
