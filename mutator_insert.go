// archivo: mutator_inser.go
package mutator

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func (m *Mutator) insertChars(runes []rune, mutationType MutationType) []rune {
	var counterInsert int
	for counterInsert < mutationType.Length {
		pos := rand.Intn(len(runes) + 1)
		runes = insertChar(runes, pos, func() rune { return m.randomChar(mutationType) })
		counterInsert++
	}
	return runes
}

func (m *Mutator) insertBoundaryValues(runes []rune) []rune {
	boundaryValues := []string{
		"0", "1", "-1", "32767", "-32768", "2147483647", "-2147483648", "9223372036854775807", "-9223372036854775808",
	}

	// Select a random boundary value
	randomBoundary := boundaryValues[rand.Intn(len(boundaryValues))]

	// Convert the boundary value to a slice of runes
	boundaryRunes := []rune(randomBoundary)

	// Select a random position in the slice of runes
	pos := rand.Intn(len(runes) + 1)

	// Insert the boundary value at the random position
	newRunes := append(runes[:pos], append(boundaryRunes, runes[pos:]...)...)

	return newRunes
}
func (m *Mutator) insertHTML(runes []rune, mutationType MutationType) []rune {
	// Define some sample HTML tags to insert
	htmlTags := []string{
		"<div>", "</div>", "<span>", "</span>", "<a href='#'>", "</a>",
		"<b>", "</b>", "<i>", "</i>", "<p>", "</p>",
		"<br>", "<img src='#'/>", "<h1>", "</h1>",
	}

	for i := 0; i < mutationType.Length; i++ {
		// Generate a random HTML tag from the list
		randomHTML := htmlTags[rand.Intn(len(htmlTags))]

		// Convert the HTML string to a slice of runes
		htmlRunes := []rune(randomHTML)

		// Select a random position in the slice of runes
		pos := rand.Intn(len(runes) + 1)

		// Insert the HTML string at the random position
		runes = append(runes[:pos], append(htmlRunes, runes[pos:]...)...)
	}

	return runes
}

func (m *Mutator) insertFakeWord(runes []rune, mutationType MutationType) []rune {
	vowels := []rune("aeiou")
	consonants := []rune("bcdfghjklmnpqrstvwxyz")

	// Generar la palabra ficticia
	var fakeWord []rune
	for i := 0; i < mutationType.Length; i++ {
		if i%2 == 0 {
			// Insertar una consonante
			randomConsonant := consonants[rand.Intn(len(consonants))]
			fakeWord = append(fakeWord, randomConsonant)
		} else {
			// Insertar una vocal
			randomVowel := vowels[rand.Intn(len(vowels))]
			fakeWord = append(fakeWord, randomVowel)
		}
	}
	// Seleccionar una posición aleatoria en el slice runes
	pos := rand.Intn(len(runes) + 1)

	// Insertar la palabra generada en la posición seleccionada
	newRunes := append(runes[:pos], append(fakeWord, runes[pos:]...)...)

	return newRunes
}

func (m *Mutator) insertRandomJSON(runes []rune, maxDepth int) []rune {
	// Generate a random JSON object
	jsonData := m.generateRandomJSON(maxDepth, 0)

	// Convert the generated JSON object to a string
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return runes
	}
	jsonString := string(jsonBytes)

	// Convert the JSON string to a slice of runes
	jsonRunes := []rune(jsonString)

	// Select a random position in the slice of runes
	pos := rand.Intn(len(runes) + 1)

	// Insert the JSON string at the random position
	newRunes := append(runes[:pos], append(jsonRunes, runes[pos:]...)...)

	return newRunes
}

// Método de Mutator para insertar una palabra aleatoria en runes
func (m *Mutator) insertRandomWord(runes []rune, mutationType MutationType) []rune {
	var words []string
	var err error

	// Cargar el diccionario si está especificado
	if mutationType.Dictionary != "" {
		words, err = loadDictionary(mutationType.Dictionary)
		if err != nil {
			fmt.Println("Error cargando el diccionario:", err)
			return runes
		}
	}

	// Si no hay palabras en la lista de palabras ni en el diccionario, retornar los runes originales
	if len(mutationType.Wordslist) == 0 && len(words) == 0 {
		return runes
	}

	// Si hay palabras en la lista de palabras, usarlas
	if len(mutationType.Wordslist) > 0 {
		words = mutationType.Wordslist
	}

	// Insertar palabras aleatorias en la posición seleccionada
	var counterInsert int
	for counterInsert < mutationType.Length {
		// Seleccionar una palabra aleatoria
		randomWord := getRandomWord(words)

		// randomRunes := []rune(randomWord + " ") // or just delete?
		randomRunes := []rune(randomWord) // or just delete?

		// Seleccionar una posición aleatoria para insertar la palabra
		pos := rand.Intn(len(runes) + 1)

		// Insertar la palabra en la posición seleccionada
		runes = append(runes[:pos], append(randomRunes, runes[pos:]...)...)
		counterInsert++
	}

	return runes
}
