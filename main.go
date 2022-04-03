package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nПривет! Это мое первое приложение на GO, простой калькулятор")
	fmt.Println("Введите действие которое вы хотите, выполнить: +, -, *, /:")
	var action string
	var first float64
	var second float64
	fmt.Scan(&action)
	fmt.Println("Введите первое число")
	fmt.Scan(&first)
	fmt.Println("Введите второе число")
	fmt.Scan(&second)

	switch {
	case action == "+":
		sum(first, second)
	case action == "-":
		minus(first, second)
	case action == "*":
		multiplication(first, second)
	case action == "/":
		if second != 0 {
			division(first, second)
		} else {
			fmt.Println("Делить на ноль не стану, можешь не просить")
		}

	default:
		fmt.Println("Калькулятор пока не настолько умный чтобы выполнять это действие")
	}
	main()
}

func sum(a float64, b float64) float64 {
	fmt.Println("Сумма чисел:" + fmt.Sprint(a+b))
	return a + b
}

func minus(a float64, b float64) float64 {
	fmt.Println("Разность чисел:" + fmt.Sprint(a-b))
	return a - b
}
func multiplication(a float64, b float64) float64 {
	fmt.Println("Произведение чисел:" + fmt.Sprint(a*b))
	return a * b
}
func division(a float64, b float64) float64 {
	fmt.Println("Xfcnyjt чисел:" + fmt.Sprint(a/b))
	return a / b
}
