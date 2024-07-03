// Archivo: dict.go
package github.com/jaavier/mutator

import (
	"math/rand"
	"strings"
	"time"
)

func isValidWord(word string, letters, numbers, special bool) bool {
	for _, r := range word {
		if (letters && isLetter(r)) || (numbers && isNumber(r)) || (special && isSpecialChar(r)) {
			return true
		}
	}
	return false
}

func CreateDict(content string, totalWords int, kind string) map[string]int {
	var minRepetitions = 5
	var minLength = 4
	words := strings.Split(content, " ")
	letters, numbers, special := processKind(kind)
	wordCount := make(map[string]int)
	filteredWords := make(map[string]int)

	// Filtrar las palabras según el tipo de caracteres especificados
	for _, word := range words {
		word = strings.ReplaceAll(word, "\n", "")
		if isValidWord(word, letters, numbers, special) {
			wordCount[word]++
		}
	}

	// Filtrar palabras que tienen al menos 5 repeticiones y longitud mínima de 4 caracteres
	for word, count := range wordCount {
		if count >= minRepetitions && len(word) >= minLength {
			filteredWords[word] = count
		}
	}

	// Si hay menos palabras filtradas que totalWords, retornar todas las palabras filtradas
	if len(filteredWords) <= totalWords {
		return filteredWords
	}

	rand.Seed(time.Now().UnixNano())
	selectedWords := make(map[string]int)
	usedIndices := make(map[int]bool)
	filteredWordsList := make([]string, 0, len(filteredWords))

	for word := range filteredWords {
		filteredWordsList = append(filteredWordsList, word)
	}

	// Seleccionar palabras aleatorias sin repetición
	for len(selectedWords) < totalWords {
		index := rand.Intn(len(filteredWordsList))
		if !usedIndices[index] {
			word := filteredWordsList[index]
			selectedWords[word] = filteredWords[word]
			usedIndices[index] = true
		}
	}

	return selectedWords
}
