package decorator

import "log"

type CustomerSaveLogger struct {
	save CustomerSaver
}

func NewCustomerSaveLog(saver CustomerSaver) CustomerSaver {
	return &CustomerSaveLogger{save: saver}
}

func (c CustomerSaveLogger) Do(data *Customer) error {
	log.Println("Save customer")
	return c.save.Do(data)
}
