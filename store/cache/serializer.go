package cache

import (
	"errors"
	"fmt"
	"github.com/c12s/oort/domain/checker"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/proto/common"
	"google.golang.org/protobuf/proto"
	"strconv"
)

type protoAttributeSerializer struct {
}

func NewProtoAttributeSerializer() checker.AttributeSerializer {
	return protoAttributeSerializer{}
}

func (s protoAttributeSerializer) Serialize(attributes []model.Attribute) ([]byte, error) {
	list := common.AttributeList{
		Attributes: make([]*common.Attribute, 0),
	}
	for _, attribute := range attributes {
		protoAttribute := &common.Attribute{
			Id: &common.AttributeId{
				Name: attribute.Name(),
			},
			Kind: common.Attribute_AttributeKind(attribute.Kind()),
		}
		switch attribute.Kind() {
		case model.Int64:
			value := common.Int64Attribute{
				Value: attribute.Value().(int64),
			}
			protoValue, err := proto.Marshal(&value)
			if err != nil {
				return nil, err
			}
			protoAttribute.Value = protoValue
			list.Attributes = append(list.Attributes, protoAttribute)
		case model.Float64:
			value := common.Float64Attribute{
				Value: attribute.Value().(float64),
			}
			protoValue, err := proto.Marshal(&value)
			if err != nil {
				return nil, err
			}
			protoAttribute.Value = protoValue
			list.Attributes = append(list.Attributes, protoAttribute)
		case model.Bool:
			value := common.BoolAttribute{
				Value: attribute.Value().(bool),
			}
			protoValue, err := proto.Marshal(&value)
			if err != nil {
				return nil, err
			}
			protoAttribute.Value = protoValue
			list.Attributes = append(list.Attributes, protoAttribute)
		case model.String:
			value := common.StringAttribute{
				Value: attribute.Value().(string),
			}
			protoValue, err := proto.Marshal(&value)
			if err != nil {
				return nil, err
			}
			protoAttribute.Value = protoValue
			list.Attributes = append(list.Attributes, protoAttribute)
		}
	}
	bytes, err := proto.Marshal(&list)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s protoAttributeSerializer) Deserialize(bytes []byte) ([]model.Attribute, error) {
	attributes := &common.AttributeList{}
	err := proto.Unmarshal(bytes, attributes)
	if err != nil {
		return nil, err
	}
	return attributes.MapToDomain()
}

type customCheckPermissionRespSerializer struct {
}

func NewCustomCheckPermissionSerializer() checker.CheckPermissionResponseSerializer {
	return customCheckPermissionRespSerializer{}
}

func (c customCheckPermissionRespSerializer) Serialize(resp checker.CheckPermissionResp) ([]byte, error) {
	allowed := "t"
	if !resp.Allowed {
		allowed = "f"
	}
	var err string
	if resp.Error != nil {
		err = resp.Error.Error()
	} else {
		err = ""
	}
	return []byte(fmt.Sprintf("%s%s", err, allowed)), nil
}

func (c customCheckPermissionRespSerializer) Deserialize(bytes []byte) (checker.CheckPermissionResp, error) {
	s := string(bytes)
	allowed, err := strconv.ParseBool(s[len(s)-1:])
	if err != nil {
		return checker.CheckPermissionResp{}, err
	}
	var respErr error
	if len(s) == 1 {
		respErr = nil
	} else {
		respErr = errors.New(s[:len(s)-1])
	}
	return checker.CheckPermissionResp{
		Allowed: allowed,
		Error:   respErr,
	}, nil
}
