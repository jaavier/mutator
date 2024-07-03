package mutator

import (
	"fmt"
	"reflect"
	"strconv"
)

func MutateModel(model interface{}) {
	val := reflect.ValueOf(model).Elem() // Obtener el valor subyacente del puntero

	if val.Kind() != reflect.Struct {
		fmt.Println("Expected a pointer to a struct")
		return
	}

	// Recorrer los campos del struct
	for i := 0; i < val.NumField(); i++ {
		var allMutations = []MutationType{}
		field := val.Field(i)
		fieldType := field.Type() // Obtener el tipo del campo
		fieldValue := field.Interface()
		var initialText string
		var mutatedValue string

		// Convertir el valor a una cadena adecuada según el tipo
		switch fieldType.Kind() {
		case reflect.String:
			initialText = fieldValue.(string)
			allMutations = append(
				allMutations,
				NewInsertMutation(GenerateRandom(1, 1000), 0.25, "all"),
				NewHtmlElement(GenerateRandom(1, 50), 0.25),
				NewJSONMutation(GenerateRandom(1, 10), 0.25),
				NewCaseChangeMutation(GenerateRandom(1, 10), 0.25),
				NewBoundaryMutation(GenerateRandom(1, 20), 0.25),
				NewPickMutation(GenerateRandom(1, 20), 0.25, "all"),
				NewSwapMutation(GenerateRandom(1, 20), 0.25),
				NewVowelChangeMutation(GenerateRandom(1, 10), 0.25),
			)
		case reflect.Int:
			initialText = strconv.Itoa(fieldValue.(int))
			allMutations = append(allMutations, NewBoundaryMutation(GenerateRandom(0, 1), 1))
			allMutations = append(allMutations, NewBoundaryMutation(GenerateRandom(0, 1), 1))
		case reflect.Float64:
			initialText = fmt.Sprintf("%f", fieldValue.(float64))
			allMutations = append(allMutations, NewBoundaryMutation(GenerateRandom(1, 10), 1))
		default:
			// Para otros tipos, conviértelos a cadena
			initialText = fmt.Sprintf("%v", fieldValue)
		}
		// Aplicar la mutación solo si initialText no está vacío
		for {
			mutatedValue = New(&Config{
				InitialText:   initialText,
				MutationTypes: allMutations,
			}).ApplySingleMutation()
			if mutatedValue != initialText {
				break
			}
		}

		// Convertir el valor mutado de vuelta al tipo original y asignarlo al campo
		switch fieldType.Kind() {
		case reflect.String:
			field.SetString(mutatedValue)
		case reflect.Int:
			mutatedInt, err := strconv.Atoi(mutatedValue)
			if err == nil {
				field.SetInt(int64(mutatedInt))
			}
		case reflect.Float64:
			mutatedFloat, err := strconv.ParseFloat(mutatedValue, 64)
			if err == nil {
				field.SetFloat(mutatedFloat)
			}
		case reflect.Bool:
			mutatedBool := RandomBool()
			field.SetBool(mutatedBool)
		default:
			// Para otros tipos, no se realiza la asignación
			fmt.Printf("Unsupported field type: %s\n", fieldType)
		}
	}
}
