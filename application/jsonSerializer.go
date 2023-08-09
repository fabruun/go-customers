package application

import (
	"encoding/json"
	"fmt"
	"os"

	domain "github.com/fabruun/go-customers/domain"
)

type JsonMapperinterface interface {
	ReadJsonFile(path string)
	MapDataFromJSON() *domain.Customers
}

type JsonMapper struct {
	JsonFile  *os.File          `json:"file"`
	ByteValue []byte            `json:"byte_value"`
	customers *domain.Customers `json:"customers"`
}

func (p *JsonMapper) ReadJsonFile(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	p.JsonFile = jsonFile
}

func (p *JsonMapper) MapDataFromJSON() *domain.Customers {
	byteValue, err := os.ReadFile(p.JsonFile.Name())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
  p.customers = 
	var customers domain.Customers
	json.Unmarshal(byteValue, &customers)


	return *customers
}
