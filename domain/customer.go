package domain

import "github.com/google/uuid"

type Customers struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Address   Address   `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}
