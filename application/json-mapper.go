package application

import (
	"encoding/json"
	"log"
	"os"

	domain "github.com/fabruun/go-customers/domain"
)

type JsonMapperinterface interface {
	ReadJsonFile(path string)
	MapDataFromJSON()
}

type JsonMapper struct {
	JsonFile  *os.File         `json:"file"`
	ByteValue []byte           `json:"byte_value"`
	Customers domain.Customers `json:"customers"`
}

func (p *JsonMapper) ReadJsonFile(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	p.JsonFile = jsonFile
}

func (p *JsonMapper) MapDataFromJSON() {
	byteValue, err := os.ReadFile(p.JsonFile.Name())
	if err != nil {
		path := p.JsonFile.Name()
		log.Fatalf("Failed to read %s. With error: %v", path, err)
	}
	var customers domain.Customers
	p.ByteValue = byteValue
	json.Unmarshal(byteValue, &customers.Customers)
	p.Customers = customers
}
