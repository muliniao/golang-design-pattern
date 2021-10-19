package sample

type DataSourceDecorator struct {
	dataSource	DataSource
}

func NewDataSourceDecorator(source DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{
		dataSource: source,
	}
}

func (f *DataSourceDecorator) WriteData(data []byte) {
	f.dataSource.WriteData(data)
}

func (f *DataSourceDecorator) ReadData() []byte {
	return f.dataSource.ReadData()
}