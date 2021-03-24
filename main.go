package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" 
	"ecommerce-backend/db"
	"ecommerce-backend/product"
)

func main() {
	fmt.Println(" ")
	fmt.Println("---------------------------")
	fmt.Println("Welcome to the backend...")
	fmt.Println("---------------------------")
	fmt.Println(" ")
	product.Connection = db.GetConnection()
	product.GetProduct(12312312)

	product.GetAllProducts()
}