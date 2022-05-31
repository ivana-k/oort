package model

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"time"
)

type Path []model.Resource

func (path Path) Path(resourceVar string) (pattern string, lastResourceVar string) {
	last := len(path) - 1
	lastResourceVar = fmt.Sprintf("%s%d", resourceVar, last)
	for i, resource := range path {
		pattern += fmt.Sprintf("(%s%d: %s{ ", resourceVar, i, resourceLabel)
		for key, value := range resource.GetArgs() {
			pattern += fmt.Sprintf("%s: %s,", key, getValue(value))
		}
		pattern = pattern[:len(pattern)-1]
		pattern += "})"
		if i != last {
			pattern += fmt.Sprintf("-[:%s]->", parentRelationship)
		}
	}
	return
}

func (path Path) IdentityPath(resourceVar string) (pattern string, lastResourceVar string) {
	last := len(path) - 1
	lastResourceVar = fmt.Sprintf("%s%d", resourceVar, last)
	for i, resource := range path {
		if i != last {
			pattern += fmt.Sprintf("(%s%d: %s{ ", resourceVar, i, resourceLabel)
		} else {
			pattern += fmt.Sprintf("(%s%d:%s:%s{ ", resourceVar, i, resourceLabel, identityLabel)
		}
		for key, value := range resource.GetArgs() {
			pattern += fmt.Sprintf("%s: %s,", key, getValue(value))
		}
		pattern = pattern[:len(pattern)-1]
		pattern += "})"
		if i != last {
			pattern += fmt.Sprintf("-[:%s]->", parentRelationship)
		}
	}
	return
}

//func (path Path) ReversedPath(resourceVar string) (pattern string, lastResourceVar string) {
//	last := len(path) - 1
//	lastResourceVar = fmt.Sprintf("%s%d", resourceVar, 0)
//	for i, resource := range reverse(path) {
//		pattern += fmt.Sprintf("(%s%d: %s{ ", resourceVar, i, resourceLabel)
//		for key, value := range resource.GetArgs() {
//			pattern += fmt.Sprintf("%s: %s,", key, getValue(value))
//		}
//		pattern = pattern[:len(pattern)-1]
//		pattern += "})"
//		if i != last {
//			pattern += fmt.Sprintf("<-[:%s]-", parentRelationship)
//		}
//	}
//	return
//}

func (path Path) ParentRelationship() string {
	return parentRelationship
}

func (path Path) PermissionRelationship() string {
	return permissionRelationship
}

func getValue(value interface{}) string {
	switch value.(type) {
	case time.Time:
		return fmt.Sprintf("datetime('%s')", value.(time.Time).Format(datetimeLayout))
	case string:
		return fmt.Sprintf("\"%s\"", value)
	default:
		return fmt.Sprint(value)
	}
}

func reverse(path Path) Path {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
