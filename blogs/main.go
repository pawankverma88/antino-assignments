package main

import (
	"data"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	var err error
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbReadHost := "localhost"
	dbReadPort := "3306"
	dbReadUser := "root"
	dbReadPass := ""

	// Connect the database
	_, err = data.InitDB(dbHost, dbPort, dbUser, dbPass, dbReadHost, dbReadPort, dbReadUser, dbReadPass)
	if err != nil {
		log.Panic(err)
	}

	router := httprouter.New()
	router.RedirectTrailingSlash = true
	addRouteHandlers(router)

	fmt.Println("Setup complete. Running API server...")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "Authorization"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	log.Fatal(http.ListenAndServe(":5000", c.Handler(router)))
}
