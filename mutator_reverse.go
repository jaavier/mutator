package github.com/jaavier/mutator

func (m *Mutator) reverseSubstring(runes []rune, mutationType MutationType) []rune {
	if len(runes) == 0 {
		return runes
	}

	length := mutationType.Length
	if length <= 0 || length > len(runes) {
		length = len(runes)
	}

	// Escoge una posición de inicio aleatoria tal que la subcadena completa pueda caber
	start := 0
	end := start + length

	// Crear un nuevo slice para almacenar los caracteres invertidos
	reversed := make([]rune, len(runes))
	copy(reversed, runes)

	// Invierte la subcadena entre start y end
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = runes[j], runes[i]
	}

	return reversed
}
