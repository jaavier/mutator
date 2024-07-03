// archivo: mutator_delete.go
package github.com/jaavier/mutator

import "math/rand"

func (m *Mutator) deleteChars(runes []rune, mutationType MutationType) []rune {
	var counterDelete int
	validDeletePositions := m.getValidDeletePositions(runes, mutationType)
	for counterDelete < mutationType.Length && len(validDeletePositions) > 0 {
		pos := validDeletePositions[rand.Intn(len(validDeletePositions))]
		runes = deleteChar(runes, pos)
		counterDelete++
		validDeletePositions = m.getValidDeletePositions(runes, mutationType)
	}
	return runes
}
