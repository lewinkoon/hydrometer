package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type ReservoirData struct {
	variable  string
	value     string
	timestamp string
}

func main() {

	var datatable []ReservoirData

	// creating a new Colly instance
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	// visiting the target page
	var url string = "https://www.saihduero.es/risr/EM171"
	c.Visit(url)

	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"variable",
		"valor",
		"fecha",
	}
	writer.Write(headers)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnHTML("table:first-of-type > tbody > tr > td", func(e *colly.HTMLElement) {
		data := ReservoirData{}
		data.variable = e.ChildText("td:nth-of-type(1)")
		fmt.Println("Body scrapping complete")
	})

	// writing each item as a CSV row
	for _, item := range datatable {
		// converting a struct to an array of strings
		record := []string{
			item.variable,
			item.value,
			item.timestamp,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()

}
