// Archivo: helpers.go
package github.com/jaavier/mutator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// changeCase cambia el caso del carácter dado.
func changeCase(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	} else if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

// duplicateChar duplica el carácter en la posición dada.
func duplicateChar(runes []rune, pos int) []rune {
	if pos < 0 || pos >= len(runes) {
		return runes
	}
	char := runes[pos]
	runes = insertChar(runes, pos+1, func() rune { return char })
	return runes
}

// containsRune verifica si una cadena contiene un carácter específico.
func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

func swapChars(runes []rune, pos1 int, pos2 int) []rune {
	runes[pos1], runes[pos2] = runes[pos2], runes[pos1]
	return runes
}

// deleteChar elimina un carácter en la posición dada.
func deleteChar(runes []rune, pos int) []rune {
	if pos < 0 || pos >= len(runes) {
		return runes
	}
	return append(runes[:pos], runes[pos+1:]...)
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSpecialChar(r rune) bool {
	if r == ' ' {
		return true
	}
	return !isLetter(r) && !isNumber(r) && !containsRune(" \t\n\r", r)
}

func insertChar(runes []rune, pos int, randomChar func() rune) []rune {
	// Asegúrate de que el índice esté en el rango válido
	if pos < 0 {
		pos = 0
	} else if pos > len(runes) {
		pos = len(runes)
	}

	char := randomChar()

	// Insertar el carácter en la posición deseada
	runes = append(runes[:pos], append([]rune{char}, runes[pos:]...)...)

	return runes
}

func GenerateRandom(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func buildCharSet(mutationType MutationType) []rune {
	var charSet []rune
	if mutationType.Letters {
		charSet = append(charSet, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")...)
	}
	if mutationType.Numbers {
		charSet = append(charSet, []rune("0123456789")...)
	}
	if mutationType.Special {
		charSet = append(charSet, []rune("!@#$%^&*()_+-=[]{}|;:',.<>/?~ ")...)
	}
	return charSet
}

func insertStaticText(text string, pos int, staticText string) string {
	runes := []rune(text)
	if pos < 0 {
		pos = 0
	} else if pos > len(runes) {
		pos = len(runes)
	}
	pre := string(runes[:pos])
	post := string(runes[pos:])
	return fmt.Sprintf("%s%s%s", pre, staticText, post)
}

// processKind analiza la cadena kind y activa los flags correspondientes
func processKind(kind string) (bool, bool, bool) {
	letters := false
	numbers := false
	special := false

	// Dividir la cadena kind por comas
	kinds := strings.Split(kind, ",")

	for _, k := range kinds {
		switch strings.TrimSpace(k) {
		case "letters":
			letters = true
		case "numbers":
			numbers = true
		case "special":
			special = true
		case "all":
			letters = true
			numbers = true
			special = true
		}
	}

	return letters, numbers, special
}

// Función para leer el diccionario y devolver una lista de palabras
func loadDictionary(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// Función para seleccionar una palabra aleatoria del diccionario
func getRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}
