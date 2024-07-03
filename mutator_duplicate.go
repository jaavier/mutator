// archivo: mutator_duplicate.go

package mutator

import "math/rand"

func (m *Mutator) duplicateChars(runes []rune, mutationType MutationType) []rune {
	var counterDuplicate int
	for counterDuplicate < mutationType.Length {
		if len(runes) > 0 {
			pos := rand.Intn(len(runes))
			runes = duplicateChar(runes, pos)
			counterDuplicate++
		}
	}
	return runes
}
