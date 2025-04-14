package fundamentos

import "fmt"

func aritmeticos() {
	fmt.Println("Soma:", 10+2)
	fmt.Println("Subtração:", 10-2)
	fmt.Println("Multiplicação:", 10*2)
	fmt.Println("Divisão:", 10/2)
	fmt.Println("Resto:", 10%2)

}

func relacionais() {
	fmt.Println("Maior que:", 10 > 2)
	fmt.Println("Menor que:", 10 < 2)
	fmt.Println("Maior ou igual:", 10 >= 2)
	fmt.Println("Menor ou igual:", 10 <= 2)
	fmt.Println("Igual:", 10 == 2)
	fmt.Println("Diferente:", 10 != 2)
}

func logicos() {
	fmt.Println("E:", true && false)
	fmt.Println("Ou:", true || false)
	fmt.Println("Não:", !true)
}
