package main

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Constructor struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func Create(c echo.Context) error {
	logFields := logrus.Fields{
		"_event":           "Create",
		"http_status_code": c.Response().Status,
	}

	constructor := &Constructor{}
	err := json.NewDecoder(c.Request().Body).Decode(constructor)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return c.JSON(500, echo.Map{
			"error": "bad request body",
		})
	}

	logFields["http_request_method"] = c.Request().Method
	logFields["http_request_body"] = constructor
	logrus.WithFields(logFields).Info("Success Execute")

	return c.JSON(200, constructor)
}

func main() {
	r := echo.New()
	r.Use(middleware.Recover())

	logrus.SetFormatter(&logrus.JSONFormatter{})
	f, err := os.OpenFile("logs/app.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Error(err)
	}
	logrus.SetOutput(f)

	r.POST("/create", Create)
	if err := r.Start(":3000"); err != nil {
		os.Exit(1)
	}
}
