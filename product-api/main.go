package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shashwatkoliwad/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	http.ListenAndServe(":8080", ph)
}
