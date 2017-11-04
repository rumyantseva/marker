package main

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"
)

// Catalog descrbes catalog.xml data
type Catalog struct {
	XMLName xml.Name `xml:"catalog"`
	Items   []Item   `xml:"item"`
}

// Item describes an item of data
type Item struct {
	Name  string `xml:"name"`
	Title string `xml:"title"`
	Hex   string `xml:"hex"`
}

func main() {
	log.Print("Start prosessing...")

	fin, err := os.Open("./catalog.xml")
	if err != nil {
		log.Fatalf("Error opening catalog file: %+v\n", err)
	}
	defer fin.Close()

	var catalog Catalog
	err = xml.NewDecoder(fin).Decode(&catalog)
	if err != nil {
		log.Fatalf("Error decoding data: %+v\n", err)
	}

	if len(catalog.Items) == 0 {
		log.Fatal("The catalog is empty.")
	}

	_, err = os.Stat("./catalog.csv")
	if os.IsNotExist(err) {
		_, err = os.Create("./catalog.csv")
	} else if err != nil {
		log.Fatalf("Error creating file: %+v\n", err)
	}

	fout, err := os.OpenFile("./catalog.csv", os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Error opening catalog file: %+v\n", err)
	}
	defer fout.Close()

	writer := csv.NewWriter(fout)
	err = writer.Write([]string{"name", "title", "hex"})
	if err != nil {
		log.Fatalf("Couldn't write data: %+v\n", err)
	}
	for _, item := range catalog.Items {
		err := writer.Write([]string{item.Name, item.Title, item.Hex})
		if err != nil {
			log.Fatalf("Couldn't write data: %+v\n", err)
		}
		writer.Flush()
	}

	log.Print("Processing is finished.")
}
