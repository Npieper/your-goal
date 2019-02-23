package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

/* type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}
*/

func home(w http.ResponseWriter, r *http.Request) {
	data, _ := GetAllExpenses()
	// data, _ := GetAllExpensesPrice()

	monthlyExpenses := calculateMonthlyExpenses(data)

	expense := Expenses{}
	expense.MonthlyExpenses = monthlyExpenses
	expense.Expense = data

	t, _ := template.ParseFiles("assets/index.html")
	t.Execute(w, expense)
}

func createExpense(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		expenseName := r.FormValue("name")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

		todo := Expense{Price: price, Name: expenseName}
		todo.test123()

	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	data, _ := GetAllExpensesPrice()
	fmt.Println(data)

	t, _ := template.ParseFiles("assets/index.html")
	t.Execute(w, data)
}

func getExpenseById(w http.ResponseWriter, r *http.Request) {

}

func calculateExpenses(w http.ResponseWriter, r *http.Request) {

}
