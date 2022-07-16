package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	e := echo.New()
	e.Use(loggingMiddleware)

	e.GET("/isInt", func(c echo.Context) error {
		a := c.QueryParam("a")

		if _, err := strconv.Atoi(a); err != nil {
			return c.String(http.StatusBadRequest, "not ok")
		}

		return c.String(http.StatusOK, "ok")
	})

	e.Start(":2001")
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()

		next := next(ctx)

		log.WithFields(log.Fields{
			"method":     ctx.Request().Method,
			"path":       ctx.Path(),
			"status":     ctx.Response().Status,
			"latency_ns": time.Since(start).Nanoseconds(),
		}).Info("request details")

		return next
	}
}
