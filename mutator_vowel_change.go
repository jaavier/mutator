// Archivo: vowels.go
package github.com/jaavier/mutator

import (
	"math/rand"
)

// Conjunto de vocales en minúsculas y mayúsculas
var lowercaseVowels = "aeiou"
var uppercaseVowels = "AEIOU"

// Función para cambiar las vocales en el texto
func (m *Mutator) vowelChange(runes []rune, mutationType MutationType) []rune {
	var counterVowelChange int
	var validPositions []int

	// Encuentra las posiciones de las vocales en la cadena
	for i, r := range runes {
		if containsRune(lowercaseVowels+uppercaseVowels, r) {
			validPositions = append(validPositions, i)
		}
	}

	// Baraja las posiciones válidas para el cambio de vocales
	rand.Shuffle(len(validPositions), func(i, j int) {
		validPositions[i], validPositions[j] = validPositions[j], validPositions[i]
	})

	// Cambia las vocales en las posiciones aleatorias
	for i := 0; i < len(validPositions) && counterVowelChange < mutationType.Length; i++ {
		pos := validPositions[i]
		oldVowel := runes[pos]

		// Genera una nueva vocal diferente de la vocal actual
		newVowel := m.randomDifferentVowel(oldVowel)
		runes[pos] = newVowel // Cambia la vocal
		counterVowelChange++
	}

	return runes
}

// Helper para seleccionar una vocal aleatoria diferente de la vocal actual
func (m *Mutator) randomDifferentVowel(currentVowel rune) rune {
	var newVowel rune
	var vowelSet string

	// Determina el conjunto de vocales en función del caso de la vocal actual
	if isLowercase(currentVowel) {
		vowelSet = lowercaseVowels
	} else {
		vowelSet = uppercaseVowels
	}

	// Genera una nueva vocal diferente de la vocal actual
	for {
		newVowel = rune(vowelSet[rand.Intn(len(vowelSet))])
		if newVowel != currentVowel {
			break
		}
	}
	return newVowel
}

// Función para verificar si un rune es minúscula
func isLowercase(r rune) bool {
	return r >= 'a' && r <= 'z'
}
