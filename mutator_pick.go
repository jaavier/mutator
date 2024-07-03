// archivo: mutator_pick.go
package github.com/jaavier/mutator

import (
	"math/rand"
)

func (m *Mutator) pickChars(runes []rune, mutationType MutationType) []rune {
	if len(runes) == 0 || mutationType.Length <= 0 {
		return runes
	}

	// Asegúrate de que Length no sea mayor que la longitud del texto inicial
	if mutationType.Length > len(runes) {
		mutationType.Length = len(runes)
	}

	// Procesar el kind para obtener los flags

	// Función para verificar si un rune es válido según los flags
	isValidRune := func(r rune) bool {
		if mutationType.Letters && isLetter(r) {
			return true
		}
		if mutationType.Numbers && isNumber(r) {
			return true
		}
		if mutationType.Special && isSpecialChar(r) {
			return true
		}
		return false
	}

	// Filtrar los runes según el kind
	var filteredRunes []rune
	for _, r := range runes {
		if isValidRune(r) {
			filteredRunes = append(filteredRunes, r)
		}
	}
	// Si no hay suficientes runes válidos, devolver los runes originales
	if len(filteredRunes) < mutationType.Length {
		return filteredRunes
	}

	// Selecciona una posición de inicio aleatoria en los runes filtrados
	start := rand.Intn(len(filteredRunes) - mutationType.Length + 1)
	end := start + mutationType.Length

	// Copia la subcadena seleccionada al final del texto
	selectedChars := filteredRunes[start:end]

	// Elimina el contenido original y reemplaza con los caracteres seleccionados
	return append([]rune{}, selectedChars...)
}
