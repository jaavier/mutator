package mutator

import (
	"math/rand"
	"time"
)

func (m *Mutator) swapChars(runes []rune, mutationType MutationType) []rune {
	rand.Seed(time.Now().UnixNano())
	var counterSwap int

	// Si la longitud de runes es menor o igual a 1, no hay nada que intercambiar
	if len(runes) <= 1 {
		return runes
	}

	for counterSwap < mutationType.Length {
		pos1 := rand.Intn(len(runes))
		pos2 := rand.Intn(len(runes))

		// Asegurarse de que las posiciones sean diferentes
		for pos1 == pos2 {
			pos2 = rand.Intn(len(runes))
		}

		runes = swapChars(runes, pos1, pos2)
		counterSwap++
	}
	return runes
}
