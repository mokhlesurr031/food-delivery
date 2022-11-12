package main

import (
	"fmt"
	"food-delivery/database"
	"food-delivery/migrations"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func main() {
	database.Init()
	migrations.Migrate()

	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", SayHello)

	http.ListenAndServe("0.0.0.0:3001", handler)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello World`)
}
