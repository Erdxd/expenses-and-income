package main

import (
	"fmt"
)

func main() {
	tr := make([]string, 0, 1)
	fmt.Println(tr)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1")
	fmt.Println(tr)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "2")
	fmt.Println(tr)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "3")
	fmt.Println(tr)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "4")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "5")
	fmt.Println(len(tr), cap(tr))

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
