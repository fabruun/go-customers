package application

import (
	"os"

	"github.com/fabruun/go-customers/domain"
)

type JsonDeserializerInterface interface {
	readJsonFile(path string) error
	deserializeFromJson(data []byte) ([]map[string]string, error)
	mapDataToStruct(data []map[string]string) ([]*domain.Customer, error)
}

type JsonMapper struct {
	File       []byte              `json:"file"`
	MappedData []map[string]string `json:"mapped_data"`
	Customers  []*domain.Customer  `json:"customers"`
}

func (p *JsonMapper) readJsonFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	p.File = file
	return nil
}
