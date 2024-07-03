// Archivo: mutator.go
package github.com/jaavier/mutator

type MutationAction int

const (
	Replace MutationAction = iota
	Insert
	Delete
	Bool
	Swap
	CaseChange
	Duplicate
	Reverse
	VowelChange
	Pick
	RandomWord
	FakeWord
	HtmlElement
	BoundaryValues
	RandomJSON
)

func New(config *Config) *Mutator {
	if config == nil {
		config = &Config{}
	}
	if config.MutationProb == 0 {
		config.MutationProb = 0.1
	}
	if len(config.MutationTypes) == 0 {
		config.MutationTypes = []MutationType{
			{Type: Replace, Any: true, MutationProb: config.MutationProb},
			{Type: Insert, Any: true, MutationProb: config.MutationProb},
			{Type: Delete, Any: true, MutationProb: config.MutationProb},
			{Type: Swap, Any: true, MutationProb: config.MutationProb},
			{Type: CaseChange, Any: true, MutationProb: config.MutationProb},
			{Type: Duplicate, Any: true, MutationProb: config.MutationProb},
			{Type: Reverse, Any: true, MutationProb: config.MutationProb},
		}
	} else {
		for i := range config.MutationTypes {
			mt := &config.MutationTypes[i]
			if !mt.Letters && !mt.Numbers && !mt.Special {
				mt.Any = true
			}
		}
	}
	if config.CharSet == "" {
		config.CharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:',.<>/?~ "
	}
	if config.CharProbabilities == nil {
		config.CharProbabilities = buildCharProbabilities(config.CharSet)
	}
	return &Mutator{
		InitialText:        config.InitialText,
		GlobalMutationProb: config.MutationProb,
		MutationTypes:      config.MutationTypes,
		GlobalCharSet:      config.CharSet,
		GlobalCharProb:     config.CharProbabilities,
		Prefix:             config.Prefix,     // Inicializa el campo Prefix
		Suffix:             config.Suffix,     // Inicializa el campo Suffix
		StaticText:         config.StaticText, // Inicializa el campo StaticText
	}
}

func buildCharProbabilities(charSet string) map[rune]float64 {
	probabilities := make(map[rune]float64)
	characters := []rune(charSet)
	count := len(characters)
	for _, char := range characters {
		probabilities[char] = 1.0 / float64(count)
	}
	return probabilities
}
