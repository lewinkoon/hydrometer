package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type ReservoirData struct {
	Timestamp string `json:"timestamp"`
	Variable  string `json:"variable"`
	Value     string `json:"value"`
	Unit      string `json:"unit"`
}

func Scrape(url string) []ReservoirData {
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

	c.OnHTML("div.table-responsive:not(.m-b-10) > table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			data := ReservoirData{}
			words := strings.Fields(h.ChildText("td:nth-of-type(2)"))

			data.Variable = h.ChildText("td:nth-of-type(1)")
			data.Value = strings.Replace(strings.Replace(words[0], ".", "", -1), ",", ".", -1)
			data.Unit = words[1]
			data.Timestamp = h.ChildText("td:nth-of-type(3) > span.hidden-sm-down")

			datatable = append(datatable, data)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(datatable, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("data.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}
	})

	c.Visit(url)

	return datatable
}
