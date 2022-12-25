package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type store struct {
	id   uint
	name string
	url  string
	time string
}

type product struct {
	id          uint
	store_id    uint
	name        string
	description string
	price       float64
}

func connect() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "mysql",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// allStores queries for all stores it's having
func allStores() ([]store, error) {
	// A stores slice to hold data from returned rows.

	rows, err := db.Query(`SELECT * FROM stores`)
	if err != nil {
		return nil, fmt.Errorf("stores: %v", err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	var stores []store

	for rows.Next() {
		var u store

		if err := rows.Scan(&u.id, &u.name, &u.url, &u.time); err != nil {
			return nil, fmt.Errorf("stores %v", err)
		}
		stores = append(stores, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("stores %v", err)
	}
	return stores, nil
}

// allProducts queries for products it's having
func allProducts() ([]product, error) {
	// A products slice to hold data from returned rows.

	rows, err := db.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, fmt.Errorf("products: %v", err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	var products []product

	for rows.Next() {
		var u product

		if err := rows.Scan(&u.id, &u.store_id, &u.name, &u.description, &u.price); err != nil {
			return nil, fmt.Errorf("products %v", err)
		}
		products = append(products, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("products %v", err)
	}
	return products, nil
}
