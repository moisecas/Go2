package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "se me cayo el pelao")
		return 0

	}

}
func parsear(entrada string) int {
	operador1, _ := strconv.Atoi(entrada)
	return operador1

}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

func main() {

	entrada := leerEntrada() //lo que nos va a devolver leerEntrada
	operador := leerEntrada()
	c := calc{} //se selecciona el struct
	c.operate(entrada, operador)
	fmt.Println(c.operate(entrada, operador))

}
