package models

import "main.go/internals/adapters/database"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	db := database.ConnectDB()
	defer db.Close()
	allProducts, err := db.Query("SELECT * FROM PRODUCTS")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}

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
