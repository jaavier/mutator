// archivo: mutator_change_case.go
package github.com/jaavier/mutator

import (
	"math/rand"
)

func (m *Mutator) changeCaseChars(runes []rune, mutationType MutationType) []rune {
	var counterCaseChange int
	validPositions := m.getValidChangeCasePositions(runes, mutationType)
	usedPositions := make(map[int]bool)

	for counterCaseChange < mutationType.Length && len(usedPositions) < len(validPositions) {
		// Selecciona una posición aleatoria
		pos := validPositions[rand.Intn(len(validPositions))]
		// Si la posición ya ha sido usada, continúa con la siguiente iteración
		if usedPositions[pos] {
			continue
		}
		// Marca la posición como usada
		usedPositions[pos] = true
		// Cambia el caso de la runa en la posición seleccionada
		runes[pos] = changeCase(runes[pos])
		counterCaseChange++
	}
	return runes
}
