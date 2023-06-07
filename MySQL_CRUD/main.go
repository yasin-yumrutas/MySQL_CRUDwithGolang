package main

import (
	"encoding/json"
	"net/http"
	"test/goP"
)

func PostProduct() {

}

// func main() {
// 	product := goP.Product{
// 		ID: 2, Title: "Aşıklar Ülkesi", Description: "Benim Hikayem", Price: 39.90,
// 	}

// 	goP.InsertProduct(product)
// }

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/product", addProductHandler)
	http.ListenAndServe(":3010", router)
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	var product goP.Product

	err := json.NewDecoder(r.Body).Decode(&product) //Unmarshal yerini tutar
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	goP.InsertProduct(product)

	w.WriteHeader(http.StatusCreated)
}
