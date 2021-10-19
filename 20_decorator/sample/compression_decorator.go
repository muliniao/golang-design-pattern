package sample

type CompressionDecorator struct {
	dataSourceDecorator	DataSourceDecorator
}

func NewCompressionDecorator(dataSourceDecorator DataSourceDecorator) *CompressionDecorator {
	return &CompressionDecorator{
		dataSourceDecorator:	dataSourceDecorator,
	}
}

func (c *CompressionDecorator) WriteData(data []byte) {
	c.dataSourceDecorator.WriteData(data)
	c.compress()
}

func (c *CompressionDecorator) ReadData() []byte {
	c.decompress()
	return c.dataSourceDecorator.ReadData()
}

func (c *CompressionDecorator) compress() {

}

func (c *CompressionDecorator) decompress() {

}