package models

type PersonaEstadoCivil struct {
	Id          int
	EstadoCivil *EstadoCivil
	Persona     *Persona
}
