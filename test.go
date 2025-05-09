package main

import (
	"fmt"
)

func main() {

	a := 0.0
	fmt.Println("введите транзакции (0 для выхода)")

	transaction := []float64{}
	for i := 0; ; i++ {

		fmt.Scan(&a)
		if a == 0 {

			break

		} else {
			transaction = append(transaction, a)

		}

	}

	balance := calculate(transaction)
	fmt.Println("ваш баланс", balance)

}
func calculate(transaction []float64) float64 {
	balance := 0.0
	for _, value := range transaction {
		balance += value
		fmt.Println(balance)

	}
	return balance

}
