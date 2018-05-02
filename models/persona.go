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

type Persona struct {
	Id              int
	PrimerNombre    string
	SegundoNombre   string
	PrimerApellido  string
	SegundoApellido string
	FechaNacimiento time.Time
	Usuario         *string
	Ente            int
	Foto            string
}

type Genero struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Activo            bool
	NumeroOrden       float64
}

type PersonaGenero struct {
	Id      int
	Genero  *Genero
	Persona *Persona
}

type EstadoCivil struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Activo            bool
	NumeroOrden       float64
}

type PersonaEstadoCivil struct {
	Id          int
	EstadoCivil *EstadoCivil
	Persona     *Persona
}
