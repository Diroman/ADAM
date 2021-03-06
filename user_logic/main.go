package main

import (
	"auth/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)

	port := os.Getenv("PORT")

	http.Handle("/", routes.Handlers())
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
