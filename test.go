package main

import "fmt"

func main() {
	a := 0

	transaction := []int{}
	for i := 0; ; i++ {
		fmt.Println("введите транзакции (n для выхода)")

		fmt.Scan(&a)
		if a == 0 {
			fmt.Println(transaction)
			break

		} else {
			transaction = append(transaction, a)
			fmt.Println(transaction)

		}

	}
}
