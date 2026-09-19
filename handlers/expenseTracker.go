package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DevZank/ExpenseTracker.git/models"
	"github.com/gorilla/mux"
)

type ExpenseHandler struct {
	DB *sql.DB
}

func NewExpensiveHandler(db *sql.DB) *ExpenseHandler {
	return &ExpenseHandler{DB: db}
}

func (expenseHandler *ExpenseHandler) ReadExpenses(writer http.ResponseWriter, request *http.Request) {
	rows, err := expenseHandler.DB.Query("SELECT * FROM expenses")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌    ERRO NO expenseHandler.DB.Query do READ")
	}

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ID, &expense.Description, &expense.Category, &expense.Value, &expense.Date)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("❌❌❌    ERRO NO rows.Scan do READ")
			return
		}

		expenses = append(expenses, expense)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(expenses)
}

func (expenseHandler *ExpenseHandler) CreateExpense(writer http.ResponseWriter, request *http.Request) {
	var expense models.Expense
	err := json.NewDecoder(request.Body).Decode(&expense)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌    ERRO NO json.NewDecoder(request.Body).Decode(&expense) do CREATE")
		panic(err)
	}

	err = expenseHandler.DB.QueryRow(
		"INSERT INTO expenses (description, category, value, date) VALUES ($1, $2, $3, $4) RETURNING id",
		expense.Description, expense.Category, expense.Value, expense.Date,
	).Scan(&expense.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌    expenseHandler.DB.QueryRow do CREATE")
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(expense)
}

func (expenseHandler *ExpenseHandler) UpdateExpense(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		fmt.Println("❌❌❌    trconv.Atoi do UPDATE")
		return
	}

	var expense models.Expense
	err = json.NewDecoder(request.Body).Decode(&expense)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		fmt.Println("❌❌❌    json.NewDecoder(request.Body).Decode(&expense) do UPDATE")
		return
	}

	result, err := expenseHandler.DB.Exec("UPDATE expenses SET description = $1, category = $2, value = $3, date = $4 WHERE id = $5", expense.Description, expense.Category, expense.Value, expense.Date, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌    expenseHandler.DB.Exec do UPDATE")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌    result.RowsAffected() do UPDATE")
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "No expense found with this ID", http.StatusNotFound)
		return
	}

	expense.ID = id
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(expense)
}

func (expenseHandler *ExpenseHandler) DeleteExpense(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := expenseHandler.DB.Exec("DELETE FROM expenses WHERE id = $1", id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌    expenseHandler.DB.Exec do DELETE")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		fmt.Println("❌❌❌   result.RowsAffected() do DELETE")
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "No expense found with this ID", http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
