package main

import "fmt"

func main() {
	var op, lan string
	var n1, n2 float32
	fmt.Println("select your language inserting in corresponding number 🏈/Escolha seu idioma selecionando o número correspondente⚽\n 1-English\n 2-Português")
	fmt.Scanln(&lan)
	if lan == "1" {
		fmt.Println("enter the desired operation (+, -, *, /): ")
		fmt.Scanln(&op)

		fmt.Println("Insirting the fisrt number: ")
		fmt.Scanln(&n1)

		fmt.Println("Insirting the second number")
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
				fmt.Println("Not is possible that divisor be zero")
			} else if n1 == 0 {
				fmt.Println("Not is possible that dividend be zero ") //a melhor coisa que aprendi fazendo esse repositório
			} else {
				fmt.Printf("%.2f %s %2f = %.2f", n1, op, n2, n1/n2)
			}
		default:
			fmt.Println("Invalid Operation")
		}
	} else if lan == "2" {
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
			} else if n1 == 0 {
				fmt.Println("não é possível que o dividendo seja zero ") //a melhor coisa que aprendi fazendo esse repositório
			} else {
				fmt.Printf("%.2f %s %2f = %.2f", n1, op, n2, n1/n2)
			}
		default:
			fmt.Println("Operação inválida.")
		}
	}
}
