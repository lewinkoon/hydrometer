package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type ReservoirData struct {
	Timestamp string `json:"timestamp"`
	Value     string `json:"value"`
}

func scrape(url string) []ReservoirData {
	var datatable []ReservoirData

	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnHTML("table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			data := ReservoirData{}
			data.Timestamp = h.ChildText("td.datefield")
			data.Value = strings.Replace(h.ChildText("td.valuefield"), ",", ".", 1)
			datatable = append(datatable, data)
		})
	})

	/* c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		 		js, err := json.MarshalIndent(datatable, "", "    ")
		   		if err != nil {
		   			log.Fatal(err)
		   		}
		   		fmt.Println("Writing data to file")
		   		if err := os.WriteFile("data.json", js, 0664); err == nil {
		   			fmt.Println("Data written to file successfully")
		   		}
	}) */

	c.Visit(url)

	return datatable
}
