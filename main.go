package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	conn "testapi/connection"
)





var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	_, err := conn.New(&conn.Config{
		Host:         "13.250.54.185",
		Port:         10023,
		User:         "postgres",
		Password:     "P@ssw0rd",
		DatabaseName: "test",
		Debug:        true,
	})
	if err != nil {
		log.Println("Postgres was failed, %v", err.Error())
	}


	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3\n")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "1323"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

	err = createProduct(&Product{Name:"Go product", Price: 222})
	if err != nil {
		log.Fatal(err)
	}

	print("create successful")
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.porducts(
			name, price)
			VALUES ($1, $2);"
			product.Name,
			product.Price,
		)
		return err 
}
