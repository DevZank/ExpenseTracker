package models

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Value       float64 `json:"value"`
	Date        string  `json:"date"`
}

const (
	TableName      = "expenses"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		description VARCHAR(150) NOT NULL,
		category VARCHAR(40) NOT NULL,
		value FLOAT NOT NULL,
		date VARCHAR(60)
	)`
)
