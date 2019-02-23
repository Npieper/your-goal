package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Expense struct {
	Id    int
	Price float64
	Name  string
}

type Expenses struct {
	Expense         []Expense
	MonthlyExpenses float64
}

// Db handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./nraboy.db")
	if err != nil {
		panic(err)
	}
}

func (expense *Expense) test123() (err error) {
	statement := "insert into expense(price, name) values($1, $2)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(expense.Price, expense.Name)

	return
}

/* Get all User */
func GetAllExpenses() (expenses []Expense, err error) {
	rows, err := Db.Query("select * from EXPENSE")

	if err != nil {
		return
	}

	for rows.Next() {
		expense := Expense{}
		err = rows.Scan(&expense.Id, &expense.Price, &expense.Name)
		expenses = append(expenses, expense)
	}
	rows.Close()
	return
}

/* Get all User */
func GetAllExpensesPrice() (expenses []float64, err error) {
	rows, err := Db.Query("select price from EXPENSE")

	if err != nil {
		return
	}

	for rows.Next() {

		var x float64
		err = rows.Scan(&x)
		expenses = append(expenses, x)
	}
	rows.Close()
	return
}
