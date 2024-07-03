package github.com/jaavier/mutator

import "math/rand"

// Random value generator based on type
func randomValue(propType int, maxDepth, currentDepth int, m *Mutator) interface{} {
	switch propType {
	case 0: // STRING
		return randomPropertyName()
	case 1: // NUMBER
		return rand.Intn(1000)
	case 2: // BOOL
		return rand.Intn(2) == 1
	case 3: // OBJECT
		if currentDepth < maxDepth {
			return m.generateRandomJSON(maxDepth, currentDepth+1)
		}
		return randomPropertyName()
	default:
		return nil
	}
}

// Generates a random JSON object
func (m *Mutator) generateRandomJSON(maxDepth, currentDepth int) map[string]interface{} {
	numProps := maxDepth // Number of properties per object (1 to 5)
	jsonObj := make(map[string]interface{})

	for i := 0; i < numProps; i++ {
		propName := randomPropertyName()
		propType := rand.Intn(4) // Type: 0-String, 1-Number, 2-Bool, 3-Object
		jsonObj[propName] = randomValue(propType, maxDepth, currentDepth, m)
	}

	return jsonObj
}

// Random property name generator
func randomPropertyName() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]rune, rand.Intn(10)+1)
	for i := range b {
		b[i] = rune(letters[rand.Intn(len(letters))])
	}
	return string(b)
}
