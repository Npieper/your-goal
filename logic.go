package main

import "fmt"

func calculateMonthlyExpenses(expenses []Expense) (monthlyExpenses float64) {

	monthlyExpenses = 0.0
	for _, element := range expenses {

		monthlyExpenses = monthlyExpenses + element.Price
	}
	fmt.Println(monthlyExpenses)
	return
}
