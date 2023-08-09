package domain

import (
	"github.com/fabruun/go-customers/application"
)

type Customers struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func (p *Customer) List(path string) Customers {
	jsonMapper := &application.JsonMapper{}
	jsonMapper.ReadJsonFile("data/customers.json")
	return jsonMapper.MapDataFromJSON()
}

func main() {
	customer := Customer{}               // Create a new customer
	customer.List("data/customers.json") // List all customers
}
