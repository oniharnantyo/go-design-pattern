package decorator

import "fmt"

type Customer struct {
	Name string
}

type CustomerSaver interface {
	Do(data *Customer) error
}

type customerSave struct{}

func NewCustomerSave() CustomerSaver {
	return &customerSave{}
}

func (c customerSave) Do(data *Customer) error {
	fmt.Println("Customer saved")

	return nil
}
