package mutator

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

// ExecError representa la información de error y salida de un comando ejecutado.
type ExecError struct {
	Code     int    // Código de salida del proceso
	Output   string // Salida estándar del proceso
	Error    string // Mensaje de error, si lo hubo
	IsSignal bool   // Indica si el error fue causado por una señal
	Signal   int    // Número de señal, si aplica
}

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano()) // Inicializa la semilla para la generación de números aleatorios
	return rand.Intn(2) == 0         // rand.Intn(2) genera un número 0 o 1, luego compara con 0
}

// runCommand ejecuta un comando y retorna un ExecError que contiene el código de salida, la salida estándar y el error.
func RunCommand(command string, args []string) ExecError {
	cmd := exec.Command(command, args...)

	// Capturar la salida estándar y la salida de error
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Ejecutar el comando
	err := cmd.Run()
	execErr := ExecError{
		Output: out.String(),
		Error:  stderr.String(),
	}

	if err != nil {
		// Si hay un error, intenta convertirlo en *exec.ExitError para obtener el código de salida
		if exitError, ok := err.(*exec.ExitError); ok {
			// Obtén el código de salida del proceso
			status := exitError.ExitCode()
			execErr.Code = status

			if status > 128 {
				// La diferencia entre el código de salida y 128 nos da el número de señal
				signal := status - 128
				execErr.IsSignal = true
				execErr.Signal = signal
			}
			execErr.Output = stderr.String()
		}
	}

	return execErr
}

// runPipedCommands ejecuta dos comandos con un pipe entre ellos y captura la salida y errores
func RunPipedCommands(cmd1 *exec.Cmd, cmd2 *exec.Cmd) ExecError {
	// Crear un pipe entre cmd1 y cmd2
	pipe, err := cmd1.StdoutPipe()
	if err != nil {
		return ExecError{
			Code:   -1,
			Output: "",
			Error:  fmt.Sprintf("Failed to create stdout pipe: %s", err),
		}
	}
	cmd2.Stdin = pipe

	// Capturar la salida estándar y la salida de error del segundo comando
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd2.Stdout = &out
	cmd2.Stderr = &stderr

	// Ejecutar los comandos
	err = cmd1.Start()
	if err != nil {
		return ExecError{
			Code:   -1,
			Output: "",
			Error:  fmt.Sprintf("Failed to start cmd1: %s", err),
		}
	}

	err = cmd2.Start()
	if err != nil {
		return ExecError{
			Code:   -1,
			Output: "",
			Error:  fmt.Sprintf("Failed to start cmd2: %s", err),
		}
	}

	err = cmd1.Wait()
	if err != nil {
		return ExecError{
			Code:   -1,
			Output: "",
			Error:  fmt.Sprintf("cmd1 failed: %s", err),
		}
	}

	err = cmd2.Wait()
	var exitCode int
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			exitCode = -1 // Indica que el error no es un ExitError
		}
	} else {
		exitCode = 0
	}

	return ExecError{
		Code:   exitCode,
		Output: out.String(),
		Error:  stderr.String(),
	}
}
