package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func Int(msg string) (int, error) {
	input := getInput(msg)
	inputInt, inputError := strconv.Atoi(input)
	return inputInt, inputError
}

func String(msg string) string {
	input := getInput(msg)
	return input
}

func Float32(msg string) (float32, error) {
	input := getInput(msg)
	inputFloat64, inputError := strconv.ParseFloat(input, 32)
	return float32(inputFloat64), inputError
}

func Float64(msg string) (float64, error) {
	input := getInput(msg)
	inputFloat64, inputError := strconv.ParseFloat(input, 64)
	return inputFloat64, inputError
}

func Bool(msg string) (bool, error) {
	input := getInput(msg)
	inputBool, inputError := strconv.ParseBool(input)
	return inputBool, inputError
}
