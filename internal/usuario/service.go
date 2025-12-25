package usuario

import (
	"log"
)

type Service interface {
	Create(nombre, apellido, correo, telefono string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s service) Create(nombre, apellido, correo, telefono string) error {
	log.Println("Servicio Create Usuario")
	return nil
}
