package sample

import "fmt"

type FileDataSource struct {
	FileName string
}

func NewFileDataSource(fileName string) *FileDataSource {
	return &FileDataSource{
		FileName: fileName,
	}
}

func (f *FileDataSource) WriteData([]byte) {
	fmt.Println("FileDataSource write date is calling")
}

func (f *FileDataSource) ReadData() []byte {
	fmt.Println("FileDataSource read date is calling")
	return []byte("aa")
}

