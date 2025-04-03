package main

import (
	"fmt"
	"time"
)

// Transaction представляет финансовую транзакцию
type Transaction struct {
	Amount       float64
	Date         time.Time
	IsSuccessful bool
}

func main() {
	now := time.Now()

	transactions := []Transaction{
		{Amount: 100, Date: now.AddDate(0, 0, -10), IsSuccessful: true},
		{Amount: 200, Date: now.AddDate(0, -2, 0), IsSuccessful: true},
		{Amount: 50, Date: now.AddDate(0, 0, -5), IsSuccessful: false},
	}

	// Фильтрация успешных транзакций
	var successful []Transaction
	for _, t := range transactions {
		if t.IsSuccessful {
			successful = append(successful, t)
		}
	}

	// Сумма транзакций за последний месяц
	var sumLastMonth float64
	monthAgo := now.AddDate(0, -1, 0)
	for _, t := range transactions {
		if t.Date.After(monthAgo) {
			sumLastMonth += t.Amount
		}
	}

	// Вывод результатов
	fmt.Printf("Successful transactions: %d\n", len(successful))
	fmt.Printf("Sum of transactions last month: %.2f\n", sumLastMonth)
}
