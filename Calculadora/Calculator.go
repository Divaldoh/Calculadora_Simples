package main

import "fmt"

func main() {
	var op string
	var n1, n2 float32

	fmt.Println("Digite a operação que deseja realizar (+, -, *, /): ")
	fmt.Scanln(&op)
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&n1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scanln(&n2)

	switch op {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", n1, n2, n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Não é possível que o divisor seja zero.")
		} if else {
			fmt.Printf("%.2f / %.2f = %.2f\n", n1, n2, n1/n2)//a melhor coisa que aprendi fazendo esse repositório
		} if else n1==0{
			fmt.Println("Não é possível que o dividendo seja zero.")
		}
	default:
		fmt.Println("Operação inválida.")
	}
}
