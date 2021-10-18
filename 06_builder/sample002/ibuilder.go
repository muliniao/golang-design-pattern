package sample002

type IResourceBuilder interface {
	SetName(name string)
	SetMaxTotal(maxTotal int)
	SetMaxIdle(maxIdle int)
	SetMinIdle(minIdle int)
	Validate() error
	GetResourceDetails() *ResourcePoolConfig
}