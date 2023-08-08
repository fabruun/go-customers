package application

type ServiceInterface interface {
	List() ([]Customer, error)
}

func (p *Customer) List() ([]Customer, error) {
	// TODO
}
