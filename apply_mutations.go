package github.com/jaavier/mutator

import (
	"math/rand"
)

func (m *Mutator) ApplyMutations(numMutations int) {
	m.Results = make([]string, numMutations)
	for i := 0; i < numMutations; i++ {
		mutatedText := m.mutate(m.InitialText)
		if m.StaticText != "" {
			pos := rand.Intn(len(mutatedText) + 1)
			mutatedText = insertStaticText(mutatedText, pos, m.StaticText)
		}
		mutatedText = m.Prefix + mutatedText + m.Suffix
		m.Results[i] = mutatedText
	}
}

func (m *Mutator) ApplySingleMutation() string {
	// Aplica una única mutación al texto inicial
	mutatedText := m.mutate(m.InitialText)

	// Inserta el texto estático si se ha especificado
	if m.StaticText != "" {
		pos := rand.Intn(len(mutatedText) + 1)
		mutatedText = insertStaticText(mutatedText, pos, m.StaticText)
	}

	// Añade el prefijo y sufijo al texto mutado
	mutatedText = m.Prefix + mutatedText + m.Suffix

	return mutatedText
}

func (m *Mutator) mutate(text string) string {
	runes := []rune(text)
	for _, mutationType := range m.MutationTypes {
		if rand.Float64() < m.getMutationProb(mutationType) {
			switch mutationType.Type {
			case Replace:
				runes = m.replaceChars(runes, mutationType)
			case Insert:
				runes = m.insertChars(runes, mutationType)
			case Delete:
				runes = m.deleteChars(runes, mutationType)
			case Swap:
				runes = m.swapChars(runes, mutationType)
			case CaseChange:
				runes = m.changeCaseChars(runes, mutationType)
			case Duplicate:
				runes = m.duplicateChars(runes, mutationType)
			case Reverse:
				runes = m.reverseSubstring(runes, mutationType)
			case VowelChange:
				runes = m.vowelChange(runes, mutationType)
			case Pick:
				runes = m.pickChars(runes, mutationType)
			case RandomWord:
				runes = m.insertRandomWord(runes, mutationType)
			case FakeWord:
				runes = m.insertFakeWord(runes, mutationType)
			case HtmlElement:
				runes = m.insertHTML(runes, mutationType)
			case BoundaryValues:
				runes = m.insertBoundaryValues(runes)
			case RandomJSON:
				runes = m.insertRandomJSON(runes, mutationType.Length)
			}
		}
	}
	return string(runes)
}
