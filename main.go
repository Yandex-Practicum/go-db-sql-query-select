package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}
	defer db.Close()

	// Выполнение SELECT-запроса с именованным параметром
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	// Считывание результата и заполнение массива sales
	for rows.Next() {
		var sale Sale
		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения строки результата: %w", err)
		}
		sales = append(sales, sale)
	}

	// Проверка на наличие ошибок после завершения итерации
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по строкам результата: %w", err)
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		log.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
