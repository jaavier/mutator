// archivo: mutator_replace.go
package github.com/jaavier/mutator

import "math/rand"

func (m *Mutator) replaceChars(runes []rune, mutationType MutationType) []rune {
	// Mapeo de caracteres válidos para reemplazo basado en mutationType
	var validChars []rune
	if mutationType.Any {
		if mutationType.CharSet != "" {
			validChars = []rune(mutationType.CharSet)
		} else {
			validChars = []rune(m.GlobalCharSet)
		}
	} else {
		if mutationType.CharSet != "" {
			validChars = []rune(mutationType.CharSet)
		} else {
			validChars = buildCharSet(mutationType)
		}
	}

	// Encontrar todas las posiciones válidas para reemplazo
	var validPositions []int
	for i, r := range runes {
		if m.shouldMutate(r, mutationType) {
			validPositions = append(validPositions, i)
		}
	}

	// Mezclar posiciones válidas para reemplazo aleatorio
	rand.Shuffle(len(validPositions), func(i, j int) {
		validPositions[i], validPositions[j] = validPositions[j], validPositions[i]
	})

	// Realizar el reemplazo en las posiciones aleatorias
	for i := 0; i < len(validPositions) && i < mutationType.Length; i++ {
		pos := validPositions[i]
		runes[pos] = validChars[rand.Intn(len(validChars))]
	}

	return runes
}
