package repository

import (
	"fmt"
	"net/http"
)

type Customer struct{}

func (c *Customer) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all customers.")
}
