


// Followed: https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://www.sohamkamani.com/golang/parsing-json/

package product

import (
	"fmt"
	// "log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" 
)

const (
    db_port = "3306"
	db = "mydb"
	db_user = "root"
	db_host = "127.0.0.0.1"
	db_password = "admin"
)

type Error struct {
	RequestCode int `json:"request_code"`
	ErrorCode int `json:"error_code"`
	ErrorName string `json:"name"`
	ErrorDescription string `json:"description"`
}

type Product struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Brand              string   `json:"brand"`
	Colour             string   `json:"colour"`
	Price              string   `json:"price"`
	Image              string   `json:"image"`
	Sizes              string 	`json:"sizes"`
	Special            bool     `json:"special"`
	SpecialPrice       string      `json:"special_price"`
	SpecialDescription string   `json:"special_description"`
	Category           string   `json:"category"`
	Length             string   `json:"length"`
	Width              string   `json:"width"`
}

// func main() {
// 	// conn, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/mydb")
// 	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db))
	
// 	if (err == nil) {
// 		fmt.Println("Err is nill")
// 		log.Print(db)
// 	} else {
// 		fmt.Println("Error!")
// 		log.Print(err.Error())	
// 	}

// 	GetProduct(12312312)
// 	defer conn.Close()
// }

func GetProduct(id int) Product {
	var product Product
	conn, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/mydb")
	
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

// Id                 string   `json:"id"`
// Name               string   `json:"name"`
// Brand              string   `json:"brand"`
// Colour             string   `json:"colour"`
// Price              string   `json:"price"`
// Image              string   `json:"image"`
// Sizes              string 	`json:"sizes"`
// Special            bool     `json:"special"`
// SpecialPrice       int      `json:"special_price"`
// SpecialDescription string   `json:"special_description"`
// Category           string   `json:"category"`
// Length             string   `json:"length"`
// Width              string   `json:"width"`