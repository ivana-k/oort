package model

import (
	"errors"
	"fmt"
)

type Resource interface {
	AddArg(key string, value interface{})
	GetArgs() map[string]interface{}
	GetArg(key string) (interface{}, error)
}

type resource struct {
	args map[string]interface{}
}

func NewResource() Resource {
	return &resource{
		args: map[string]interface{}{},
	}
}

func (r *resource) AddArg(key string, value interface{}) {
	r.args[key] = value
}

func (r *resource) GetArgs() map[string]interface{} {
	return r.args
}

func (r *resource) GetArg(key string) (interface{}, error) {
	value, ok := r.args[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no arg with key %s", key))
	}
	return value, nil
}
