package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(n echo.Context) error {
		datatable := Scrape("https://www.saihduero.es/risr/EM171")
		return n.JSON(http.StatusOK, datatable)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
