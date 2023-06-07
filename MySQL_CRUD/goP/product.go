package goP

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/sphinxql"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func Hata(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:321654@tcp(localhost:3306)/forgolang")
	Hata(err)

	// db.SetConnMaxIdleTime()
	// db.SetConnMaxLifetime()
	// db.SetMaxIdleConns()
	// db.SetMaxOpenConns()
}

func InsertProduct(data Product) { //---->POST
	result, err := db.Exec("INSERT INTO test(id,title,description,price)VALUES(?,?,?,?)", data.ID, data.Title, data.Description, data.Price)
	Hata(err)
	rowsAffected, err := result.RowsAffected()
	Hata(err)
	fmt.Printf("Insert'den etkilenen kayıt sayısı: %d\n", rowsAffected)
}

func UpdateProduct(data Product) { //---->PUT
	result, err := db.Exec("UPDATE test SET title=? WHERE id=?", data.Title, data.ID)
	Hata(err)
	rowsAffected, err := result.RowsAffected()
	Hata(err)
	fmt.Printf("Update'de etkilenen kayıt sayısı: %d\n", rowsAffected)
}

func DeleteProduct(data Product) { //---->DELETE
	result, err := db.Exec("DELETE FROM test WHERE id=?", data.ID)
	Hata(err)
	rowsAffected, err := result.RowsAffected()
	Hata(err)
	fmt.Printf("Delete'den etkilenen kayıt sayısı: %d\n", rowsAffected)
}

func GetProduct(data Product) { //---->SELECT
	rows, err := db.Query("SELECT * FROM test")
	Hata(err)
	if err == sql.ErrNoRows {
		fmt.Println("Hiç bir paratmetre bulanamadı")
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err = rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)
		Hata(err)
		products = append(products, prd)
	}
	for _, value := range products {
		fmt.Printf("* Select sonucu gelen veri: \n%d\n%s\n%s\n%v\n", value.ID, value.Title, value.Description, value.Price)
	}
}

func GetProductById(id int) { //---->SELECT
	var product string
	err := db.QueryRow("SELECT title FROM test WHERE id=?", id).Scan(&product)
	Hata(err)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("Bu id'de bir ürün bulunmamakta")
	default:
		fmt.Printf("Seçilen ürürün adı: %s\n", product)
	}
}
