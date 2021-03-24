package product

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" 
	// "reflect"
)

var Connection *sql.DB;


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

func GetProduct(id int) Product {
	var product Product
	// conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db))
	conn, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	
	// Execute the query
	err = conn.QueryRow("SELECT * FROM product where id = ?", id).Scan(
		&product.Id, 
		&product.Name, 
		&product.Brand, 
		&product.Colour, 
		&product.Price, 
		&product.Image, 
		&product.Sizes, 
		&product.Special, 
		&product.SpecialPrice, 
		&product.SpecialDescription, 
		&product.Category, 
		&product.Length, 
		&product.Width, 
	)
	// .Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(product.Id)
	fmt.Println(product.Name)

	return product
}

func GetAllProducts() {
	conn, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/mydb")
	ProductList := make([]*Product, 0)
	fmt.Println("GetAllProducts")
    
    // Execute the query
    results, err := conn.Query("SELECT * FROM product")
	if err != nil {
        log.Fatal(err)
    }
	
	fmt.Println("After First Error check")

	fmt.Println("Loop starting..")
	for results.Next() {
		product := new(Product)

		if err := results.Scan(
				&product.Id, 
				&product.Name, 
				&product.Brand, 
				&product.Colour, 
				&product.Price, 
				&product.Image, 
				&product.Sizes, 
				&product.Special, 
				&product.SpecialPrice, 
				&product.SpecialDescription, 
				&product.Category, 
				&product.Length, 
				&product.Width, 
		); err != nil {
			panic(err)
		}
		ProductList = append(ProductList, product)
	}

	fmt.Println("After Loop")
	if err := results.Err(); err != nil {
		panic(err)
	}

	for _, p := range ProductList {
		fmt.Println(p.Name)
	}

	defer Connection.Close()
}