package main

import (
	"fmt"
	"log"
)

func main() {
	connect()

	stores, err := allStores()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Stores found: %v\n", stores)

	products, err := allProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Stores found: %v\n", products)

	writeXml(parseXml(&stores, &products))
}
