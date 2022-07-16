package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/isInt", func(c echo.Context) error {
		a := c.QueryParam("a")

		if _, err := strconv.Atoi(a); err != nil {
			return c.String(http.StatusBadRequest, "not ok")
		}

		return c.String(http.StatusOK, "ok")
	})

	e.Start(":2001")
}
