package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	// CONECTING TO DB
	connStr := "postgres://postgres:root@localhost:5433/gopg?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// CREATE TABLE
	createProductTable(db)

	// INSERT PRODUCT AND GET ID
	product := Product{"Book", 15.15, true}
	pk := insertProcut(db, product)

	fmt.Printf("ID = %d\n", pk)

	//SELECT PRODUCT BY ID
	var name string
	var available bool
	var price float64

	query := "SELECT name, available, price FROM product WHERE id = $1"
	err = db.QueryRow(query, pk).Scan(&name, &available, &price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows found with ID")
		}
		log.Fatal(err)
	}

	fmt.Printf("Name = %s\n", name)
	fmt.Printf("Available = %t\n", available)
	fmt.Printf("Price = %f\n\n", price)

	// Query for multiple rows
	data := []Product{}
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price float64
		var available bool
		var created time.Time

		err := rows.Scan(&id, &name, &price, &available, &created)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}

	fmt.Println(data)
}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE  IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT null,
		price NUMERIC(6,2) NOT null,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProcut(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
		VALUES ($1,$2,$3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}

	return pk
}
