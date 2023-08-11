package application

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/fabruun/go-customers/domain"
)

type Customer struct{}

func (p *Customer) List(w http.ResponseWriter, req *http.Request) {
	mapper := JsonMapper{}
	mapper.ReadJsonFile(domain.JSON_PATH)
	mapper.MapDataFromJSON()
	GenerateUUID(&mapper)
	customers := mapper.Customers.Customers
	sortCustomers(&mapper)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(&customers)
	if err != nil {
		log.Fatalf("Couldn't JSON marshal response. With error: %v", err)
	}
	w.Write(response)
}

func sortCustomers(js *JsonMapper) {
	sort.SliceStable(js.Customers.Customers, func(i, j int) bool {
		return js.Customers.Customers[i].LastName < js.Customers.Customers[j].LastName
	})
}
