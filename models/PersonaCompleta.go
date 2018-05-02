package models

import (
	"time"
)

type PersonaCompleta struct {
	Id              int
	PrimerNombre    string
	SegundoNombre   string
	PrimerApellido  string
	SegundoApellido string
	FechaNacimiento time.Time
	Usuario         *string
	Ente            int
	Foto            string
	Genero          *Genero
	EstadoCivil     *EstadoCivil
}
