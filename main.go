package main

import (
	"database/sql"
	"fmt"

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

	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()

	if err != nil {
		return sales, err
	}

	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :id", sql.Named("id", client))

	if err != nil {
		return sales, err
	}

	defer rows.Close()

	for rows.Next() {
		newSale := Sale{}

		err := rows.Scan(&newSale.Product, &newSale.Volume, &newSale.Date)

		if err != nil {
			return sales, err
		}

		sales = append(sales, newSale)
	}

	if err := rows.Err(); err != nil {
		return sales, err
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
