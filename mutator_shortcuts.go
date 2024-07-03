// Archivo: easy.go
package github.com/jaavier/mutator

import (
	"math/rand"
)

// Función para crear un MutationType para un cambio de caso
func NewCaseChangeMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         CaseChange,
		Length:       length,
		MutationProb: prob,
		Letters:      true,
	}
}

// Función para crear un MutationType para un cambio de vocales
func NewVowelChangeMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         VowelChange,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para un palabras falsas
func NewFakeWordMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         FakeWord,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para inserciones
func NewInsertMutation(length int, prob float64, kind string) MutationType {
	letters, numbers, special := processKind(kind)
	return MutationType{
		Type:         Insert,
		Length:       length,
		MutationProb: prob,
		Letters:      letters,
		Numbers:      numbers,
		Special:      special,
	}
}

// Función para crear un MutationType para inserciones
func NewDictionaryWord(length int, prob float64, dictionary string, wordslist []string) MutationType {
	return MutationType{
		Type:         RandomWord,
		Length:       length,
		MutationProb: prob,
		Dictionary:   dictionary,
		Wordslist:    wordslist,
	}
}

// Función para crear un MutationType para reemplazos
func NewReplaceMutation(length int, prob float64, kind string) MutationType {
	letters, numbers, special := processKind(kind)
	return MutationType{
		Type:         Replace,
		Length:       length,
		MutationProb: prob,
		Letters:      letters,
		Numbers:      numbers,
		Special:      special,
	}
}

// Función para crear un MutationType para eliminaciones
func NewDeleteMutation(length int, prob float64, kind string) MutationType {
	letters, numbers, special := processKind(kind)
	return MutationType{
		Type:         Delete,
		Length:       length,
		MutationProb: prob,
		Letters:      letters,
		Numbers:      numbers,
		Special:      special,
	}
}

// Función para crear un MutationType para intercambios
func NewSwapMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         Swap,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para duplicaciones
func NewDuplicateMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         Duplicate,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para inversión
func NewReverseMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         Reverse,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para selección de caracteres
func NewHtmlElement(length int, prob float64) MutationType {
	return MutationType{
		Type:         HtmlElement,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para selección de caracteres
func NewPickMutation(length int, prob float64, kind string) MutationType {
	letters, numbers, special := processKind(kind)
	return MutationType{
		Type:         Pick,
		Length:       length,
		MutationProb: prob,
		Letters:      letters,
		Numbers:      numbers,
		Special:      special,
	}
}

// Función para crear un MutationType para testear limites de numeros
func NewBoundaryMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         BoundaryValues,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para testear limites de numeros
func NewBoolMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         Bool,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un MutationType para generar random jsons
func NewJSONMutation(length int, prob float64) MutationType {
	return MutationType{
		Type:         RandomJSON,
		Length:       length,
		MutationProb: prob,
	}
}

// Función para crear un Config con configuraciones predeterminadas
func NewDefaultConfig(initialText string, mutations []MutationType) *Config {
	return &Config{
		InitialText:       initialText,
		MutationTypes:     mutations,
		CharSet:           "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()",
		CharProbabilities: nil,
		Prefix:            "",
		Suffix:            "",
		StaticText:        "",
	}
}

// Function to create a random MutationType with a random kind
func NewRandomMutation(length int, prob float64) MutationType {
	// Define possible kinds
	// kinds := []string{"letters", "numbers", "special", "all"}
	// Randomly select a kind
	randomKind := "all"

	letters, numbers, special := processKind(randomKind)
	actions := []MutationAction{
		Replace,
		Insert,
		Delete,
		Swap,
		CaseChange,
		Duplicate,
		Reverse,
		VowelChange,
		Pick,
		RandomWord,
		FakeWord,
		BoundaryValues,
	}

	// Randomly select a mutation action
	randomAction := actions[rand.Intn(len(actions))]
	return MutationType{
		Type:         randomAction,
		Length:       length,
		MutationProb: prob,
		Letters:      letters,
		Numbers:      numbers,
		Special:      special,
	}
}
