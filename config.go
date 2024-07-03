// Archivo: models.go
package mutator

type MutationType struct {
	Type              MutationAction
	Letters           bool
	Numbers           bool
	Special           bool
	Any               bool
	CharSet           string
	CharProbabilities map[rune]float64
	MutationProb      float64
	Length            int
	Dictionary        string
	Wordslist         []string
}

type Config struct {
	InitialText       string
	MutationProb      float64
	MutationTypes     []MutationType
	CharSet           string
	CharProbabilities map[rune]float64
	Prefix            string // Campo para el prefijo
	Suffix            string // Campo para el sufijo
	StaticText        string // Campo para el texto estático
}

type Mutator struct {
	InitialText        string
	Results            []string
	GlobalMutationProb float64
	MutationTypes      []MutationType
	GlobalCharSet      string
	GlobalCharProb     map[rune]float64
	Prefix             string // Campo para el prefijo
	Suffix             string // Campo para el sufijo
	StaticText         string // Campo para el texto estático
}
