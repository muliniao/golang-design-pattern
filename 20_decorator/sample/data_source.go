package sample

type DataSource interface {
	WriteData([]byte)
	ReadData() []byte
}
