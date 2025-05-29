package product

import (
	"ecommerce-backend-go/db"
	"ecommerce-backend-go/queries/product"
)

func CreateProduct(name string, price float64, stock int) error {
	_, err := db.DB.Exec(product.InsertProduct, name, price, stock)
	return err
}
