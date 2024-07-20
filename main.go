package main

import (
	"fmt"
	"database/sql"
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
	if err != nil{
		return sales, err
	}
	defer db.Close()

	row, err := db.Query("select Product, Volume, Date from sales where client = :client", sql.Named("client", client))
	if err != nil {
		return sales, err
	}

	for row.Next(){
		var s Sale

		err := row.Scan(&s.Product, &s.Volume, &s.Date)
		if err != nil {
			return sales, err
		}

		sales = append(sales, s)

	}
	
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
