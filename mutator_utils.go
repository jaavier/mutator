package github.com/jaavier/mutator

// getValidDeletePositions returns valid positions for deletion based on the mutation type
func (m *Mutator) getValidDeletePositions(runes []rune, mutationType MutationType) []int {
	var positions []int
	for i, r := range runes {
		if m.shouldMutate(r, mutationType) {
			positions = append(positions, i)
		}
	}
	return positions
}

// getMutationProb returns the probability of mutation for the given mutation type
func (m *Mutator) getMutationProb(mutationType MutationType) float64 {
	if mutationType.MutationProb > 0 {
		return mutationType.MutationProb
	}
	return m.GlobalMutationProb
}

// shouldMutate checks if a rune should be mutated based on the mutation type
func (m *Mutator) shouldMutate(r rune, mutationType MutationType) bool {
	if mutationType.Any {
		return true
	}
	if mutationType.Letters && isLetter(r) {
		return true
	}
	if mutationType.Numbers && isNumber(r) {
		return true
	}
	if mutationType.Special && isSpecialChar(r) {
		return true
	}
	return false
}

// getValidChangeCasePositions returns valid positions for case change based on the mutation type
func (m *Mutator) getValidChangeCasePositions(runes []rune, mutationType MutationType) []int {
	var positions []int
	for i, r := range runes {
		if m.shouldMutate(r, mutationType) && isLetter(r) {
			positions = append(positions, i)
		}
	}
	return positions
}
