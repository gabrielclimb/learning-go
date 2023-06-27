package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"main.go/internals/models"
)

type Database struct {
	DB *sql.DB
}

func ConnectDB() *Database {
	connection := "user=postgres dbname=store password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return &Database{DB: db}
}

func (d *Database) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	allProducts, err := d.DB.Query("SELECT * FROM PRODUCTS")
	if err != nil {
		panic(err.Error())
	}
	p := models.Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	return products, err
}
