package application

import "github.com/google/uuid"

func GenerateUUID(p *JsonMapper) {
	for i := 0; i < len(p.Customers.Customers); i++ {
		p.Customers.Customers[i].ID = uuid.New()
	}
}
