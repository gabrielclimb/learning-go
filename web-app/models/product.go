package models

import (
	"main.go/internals/adapters/database"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Amount      int
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
		var id, amount int
		var name, description string
		var price float64

		err := allProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.ID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	return products, err
}

func AddNewProduct(product Product) {
	db := database.ConnectDB()
	defer db.Close()

	insertProduct, err := db.Prepare("INSERT INTO PRODUCTS (NAME, DESCRIPTION, PRICE, AMOUNT) VALUES ($1, $2, $3, $4);")
	if err != nil {
		panic(err)
	}
	_, err = insertProduct.Exec(product.Name, product.Description, product.Price, product.Amount)
	if err != nil {
		panic(err)
	}
}

func DeleteProduct(id string) {
	db := database.ConnectDB()
	defer db.Close()

	deleteProduct, err := db.Prepare("DELETE FROM PRODUCTS WHERE ID=$1")
	if err != nil {
		panic(err)
	}
	_, err = deleteProduct.Exec(id)
	if err != nil {
		panic(err)
	}
}

func EditProduct(id string) Product {
	db := database.ConnectDB()
	defer db.Close()

	editProduct, err := db.Query("SELECT * FROM PRODUCTS WHERE ID=$1;", id)
	if err != nil {
		panic(err.Error())
	}

	updateProduct := Product{}

	for editProduct.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = editProduct.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		updateProduct.Name = name
		updateProduct.Description = description
		updateProduct.Price = price
		updateProduct.Amount = amount
	}

	return updateProduct
}
