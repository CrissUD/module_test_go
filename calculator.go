package module_test_go

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Calc struct{}

func (Calc) Operate(firstValue int, secondValue int, operator string) int {
	switch operator {
	case "+":
		return firstValue + secondValue
	case "-":
		return firstValue - secondValue
	case "*":
		return firstValue * secondValue
	case "/":
		return firstValue / secondValue
	default:
		fmt.Println("Error de sintaxis con operador")
		return 0
	}
}

func (Calc) toInt(entry string) int {
	intValue, err := strconv.Atoi(entry)
	if err == nil {
		return intValue
	}
	fmt.Println("Error de sintaxis con numero")
	return 0
}

// ReadIntoKey is...
func ReadIntoKey() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
