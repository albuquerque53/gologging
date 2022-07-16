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
	log.SetFormatter(&log.JSONFormatter{})

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))

	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
}

func main() {
	e := echo.New()
	e.Use(loggingMiddleware)

	e.GET("/isInt", func(c echo.Context) error {
		value := c.QueryParam("value")

		logCtx := log.WithField("value", value)
		logCtx.Debug("parsing value")

		if _, err := strconv.Atoi(value); err != nil {
			log.WithField("value", value).Error("value is not an integer")

			return c.String(http.StatusBadRequest, "not ok")
		}

		logCtx.Debug("integer parsed")

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
