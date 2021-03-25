package product

import (
	// "fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" 
	"reflect"
	"encoding/json"
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

	// Execute the query
	err := Connection.QueryRow("SELECT * FROM product where id = ?", id).Scan(
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

	log.Println("GetProduct(): ")
	log.Println(reflect.TypeOf(product))

	return product
}

func GetAllProducts() []byte {
	
	//Declare Product List
	var ProductList []*Product 

	//Logging
	log.Println("GetAllProducts() selected")
    
    // Execute the query
    results, err := Connection.Query("SELECT * FROM product")
	if err != nil {
        log.Fatal(err)
    }

	//Lopping through results and appending to ProductList
	for results.Next() {
		product := new(Product)
		if err := results.Scan(&product.Id, 
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

		//Needs to be appending as JSON...
		ProductList = append(ProductList, product)
	}

	p, _ := json.MarshalIndent(ProductList, "", "  ")
	log.Println(string(p))

	if err := results.Err(); err != nil {
		panic(err)
	}


	//Logging
	//For Debugging: Loop through ProductList and print p.Name
	log.Println("#############################")
	log.Println(" ")
	log.Println("   Result of Database Query  ")
	log.Println(" ")
	log.Println("#############################")

	defer Connection.Close()

	return p
}