package main

import (
	"net/http"

	"github.com/DevZank/ExpenseTracker.git/config"
	"github.com/DevZank/ExpenseTracker.git/handlers"
	"github.com/DevZank/ExpenseTracker.git/models"
	"github.com/gorilla/mux"
)

func main() {
	dbConn := config.SetupDB()

	_, err := dbConn.Exec(models.CreateTableSQL)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	expenseHandler := handlers.NewExpensiveHandler(dbConn)

	router.HandleFunc("/expenses", expenseHandler.ReadExpenses).Methods("GET")
	router.HandleFunc("/expenses", expenseHandler.CreateExpense).Methods("POST")
	router.HandleFunc("/expenses/{id}", expenseHandler.UpdateExpense).Methods("PUT")
	router.HandleFunc("/expenses/{id}", expenseHandler.DeleteExpense).Methods("DELETE")

	defer dbConn.Close()

	panic(http.ListenAndServe(":8080", router))
}
