package github.com/jaavier/mutator

import (
	"math/rand"
)

// randomChar devuelve un carácter aleatorio basado en la configuración de probabilidades.
func (m *Mutator) randomChar(mutationType MutationType) rune {
	var charSet string

	// Construir charSet basado en las configuraciones de MutationType
	if mutationType.Any {
		charSet = m.GlobalCharSet
	} else {
		if mutationType.Letters {
			charSet += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if mutationType.Numbers {
			charSet += "0123456789"
		}
		if mutationType.Special {
			charSet += "!@#$%^&*()_+-=[]{}|;:',.<>/?~ "
		}
		if mutationType.CharSet != "" {
			charSet = mutationType.CharSet
		}
	}

	charProbabilities := mutationType.CharProbabilities
	if len(charProbabilities) == 0 {
		charProbabilities = m.GlobalCharProb
	}

	var chars []rune
	for char, prob := range charProbabilities {
		if rand.Float64() < prob && containsRune(charSet, char) {
			chars = append(chars, char)
		}
	}

	if len(chars) > 0 {
		return chars[rand.Intn(len(chars))]
	}

	if len(charSet) > 0 {
		return rune(charSet[rand.Intn(len(charSet))])
	}

	// Default fallback: use the global character set
	return rune(m.GlobalCharSet[rand.Intn(len(m.GlobalCharSet))])
}
