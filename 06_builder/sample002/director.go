package sample002

type Director struct {
	builder	IResourceBuilder
}

func NewDirector (builder IResourceBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() (*ResourcePoolConfig, error) {

	d.builder.SetMaxTotal(10)
	d.builder.SetMaxIdle(11)
	d.builder.SetName("Thread pool")
	d.builder.SetMinIdle(1)
	err := d.builder.Validate()
	if err != nil{
		return nil, err
	}
	return d.builder.GetResourceDetails(), nil
}