package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" 
	"ecommerce-backend/db"
	"ecommerce-backend/product"
)


type Product struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Brand              string   `json:"brand"`
	Colour             string   `json:"colour"`
	Price              string   `json:"price"`
	Image              string   `json:"image"`
	Sizes              string 	`json:"sizes"`
	Special            bool     `json:"special"`
	SpecialPrice       string   `json:"special_price"`
	SpecialDescription string   `json:"special_description"`
	Category           string   `json:"category"`
	Length             string   `json:"length"`
	Width              string   `json:"width"`
}


// This main class can be used for debugging purposes if you add a command you may want to debug into it
func main() {
	fmt.Println(" ")
	fmt.Println("-----------------------------")
	fmt.Println("   Welcome to the backend..  ")
	fmt.Println("-----------------------------")
	fmt.Println(" ")
	product.Connection = db.GetConnection()
	product.GetProduct(12312312)
	
	product.GetAllProducts()
}