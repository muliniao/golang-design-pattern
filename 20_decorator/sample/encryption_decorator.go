package sample

type EncryptionDecorator struct {
	dataSourceDecorator DataSourceDecorator
}

func NewEncryptionDecorator(dataSourceDecorator DataSourceDecorator) *EncryptionDecorator {
	return &EncryptionDecorator{
		dataSourceDecorator:	dataSourceDecorator,
	}
}

func (e *EncryptionDecorator) WriteData(data []byte) {
	e.dataSourceDecorator.WriteData(data)
	e.encryption()
}

func (e *EncryptionDecorator) ReadData() []byte {
	e.decryption()
	return e.dataSourceDecorator.ReadData()
}

func (e *EncryptionDecorator) encryption() {

}

func (e *EncryptionDecorator) decryption() {

}