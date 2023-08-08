package jsonmapper

import (
	"encoding/json"
	"os"
)

type JsonMapperInterface interface {
	ReadJsonFile(p string) ([]byte, error)
	JsonEncodeFile(f *[]byte)
}

type JsonMapper struct {
	Path   string   `json:"path"`
	Keys   []string `json:"keys"`
	Values []string `json:"values"`
}

func (r JsonMapper) ReadJsonFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (r JsonMapper) JsonEncodeFile(f *[]byte) []map[string]string {
	jsonFile, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	return []map[string]string{}
}
