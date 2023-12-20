package main

import (
	"fmt"
	"net/http"

	"app/config"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

// type Response struct {
// 	Message string `json: "message"`
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	var cfg config.EnvConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	db, err := sqlx.Connect(cfg.DBDriver, cfg.DataSourceName())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world...")

		// res := Response{Message: "hello world.../"}
		// return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)))
}
