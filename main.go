package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int    `db:"product"`
	Volume  int    `db:"volume"`
	Date    string `db:"date"`
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return []Sale{}, err
	}
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		sale := Sale{}

		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	defer db.Close()
	// напишите код здесь

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
