package sample002

import (
	"errors"
)

type ResourceBuilder struct {
	Resource	*ResourcePoolConfig
}

func NewResourceBuilder() *ResourceBuilder {
	return &ResourceBuilder{
		Resource: new(ResourcePoolConfig),
	}
}

func (r *ResourceBuilder) SetName(name string) {
	r.Resource.Name = name
}

func (r *ResourceBuilder) SetMaxTotal(maxTotal int) {
	r.Resource.MaxTotal = maxTotal
}

func (r *ResourceBuilder) SetMaxIdle(maxIdle int) {
	r.Resource.MaxIdle = maxIdle
}

func (r *ResourceBuilder) SetMinIdle(minIdle int) {
	r.Resource.MinIdle = minIdle
}

func (r *ResourceBuilder) Validate() error {

	if err := r.validateName(); err != nil {
		return err
	}

	if err := r.validateIdle(); err != nil {
		return err
	}

	return nil
}

func (r *ResourceBuilder) validateName() error  {

	if r.Resource.Name == "" {
		return errors.New("name cannot be nil")
	}

	return nil
}

func (r *ResourceBuilder) validateIdle() error  {

	if r.Resource.MinIdle > r.Resource.MaxIdle {
		return errors.New("minIdle cannot be less than maxIdle")
	}

	return nil
}

func (r *ResourceBuilder) GetResourceDetails() *ResourcePoolConfig {
	return r.Resource
}
