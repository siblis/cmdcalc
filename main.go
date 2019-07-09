package main

import (
	"cmdcalc/calculator"
	"fmt"
)

func main() {
	input := ""
	fmt.Println("Введите exit для выхода или help для примера использования")
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "help" {
			fmt.Println("Примеры:\n 2+3<Enter>\n 1+cos(Pi) <Enter>\n 2*3-arcin(0) <Enter>\n exit<Enter>")
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
