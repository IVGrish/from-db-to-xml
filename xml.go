package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Shop struct {
	XMLName      xml.Name  `xml:"shop"`
	Id           int       `xml:"id,attr"`
	Name         string    `xml:"name"`
	Url          string    `xml:"url"`
	Working_time *Time     `xml:"working_time"`
	Offers       *Products `xml:"offers"`
}

type Product struct {
	XMLName     xml.Name `xml:"item"`
	Id          int      `xml:"id,attr"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Price       float64  `xml:"price"`
}

type Time struct {
	Open  string `xml:"open"`
	Close string `xml:"close"`
}

type Products struct {
	Products []*Product `xml:"item"`
}

func parseXml(store *[]store, product *[]product) []byte {

	shop1 := &Shop{Id: 27, Name: "Coffee", Url: "shobiret.by"}
	shop1.Working_time = &Time{"10:00", "23:00"}

	var prods []*Product

	var shops []*Shop

	var ids []uint

	for _, v := range *product {
		r, _ := regexp.Compile("(<(/?[^>]+)>)")
		out := r.ReplaceAllString(string(v.description), "")

		prod := &Product{Id: int(v.id), Name: string(v.name),
			Description: string(out), Price: float64(v.price)}
		ids = append(ids, v.store_id)
		prods = append(prods, prod)
	}

	for _, v := range *store {
		shop := &Shop{Id: int(v.id), Name: string(v.name), Url: string(v.url)}
		words := strings.Split(string(v.time), "-")
		shop.Working_time = &Time{words[0], words[1]}

		add := &Products{}
		var addProduct []*Product

		for clue, item := range ids {
			if v.id == item {
				addProduct = append(addProduct, prods[clue])
			}
		}

		add.Products = addProduct
		shop.Offers = add

		shops = append(shops, shop)
	}

	nesting := shops

	out, _ := xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
	return out
}

func writeXml(out []byte) {
	fileName := "./data.xml"

	createXml := func() {
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		n, err := f.Write(out)
		if err != nil {
			panic(err)
		}
		fmt.Printf("wrote %d bytes\n", n)

		f.Sync()
	}

	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)
		if err != nil {
			panic(err)
		}
		createXml()
	} else if errors.Is(err, os.ErrNotExist) {
		createXml()
	} else {
		if err != nil {
			panic(err)
		}
	}
}
