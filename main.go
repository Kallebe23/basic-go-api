package main

import (
	"encoding/json"
	"gocommerce/config"
	"log"
	"net/http"
)

type Product struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodGet) {
		rows, err := config.DB.Query("select * from products")
		if (err != nil) {
			log.Fatal(err)
		}
		defer rows.Close()
	
		var products []Product
	
		for rows.Next() {
			var p Product
			err := rows.Scan(&p.Name, &p.Description, &p.Price)
			if (err != nil) {
				log.Fatal(err)
			}
			products = append(products, p)
		}
	
		err = rows.Err()
		if (err != nil) {
			log.Fatal(err)
		}
	
		json.NewEncoder(w).Encode(products);
	}

	if (r.Method == http.MethodPost) {
		var product Product
		json.NewDecoder(r.Body).Decode(&product)

		rows, err := config.DB.Query("insert into products values ($1, $2, $3)", product.Name, product.Description, product.Price)
		if (err != nil) {
			log.Fatal(err)
		}
		defer rows.Close()
	}

}

func main() {
	config.BootstrapDB()

	http.HandleFunc("/products", productsHandler)

	http.ListenAndServe(":3000", nil)
}