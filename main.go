package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/prabalesh/puppet/internal/db"
	"github.com/prabalesh/puppet/internal/handler"
)

var database *sql.DB

func main() {
	// database connection
	var err error
	database, err = db.InitDB("file:storage/puppet.db?cache=shared&_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}
	defer database.Close()

	mux := http.NewServeMux()

	h := handler.New(database)

	// handlers
	mux.HandleFunc("POST /api/languages", h.AddLanguage)
	mux.HandleFunc("GET /api/languages", h.ListLanguages)
	mux.HandleFunc("DELETE /api/languages/{id}", h.DeleteLanguage)
	mux.HandleFunc("POST /api/languages/{id}/installations", h.InstallLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}/installations", h.UninstallLanguage)

	// server listenting
	fmt.Println("Server started at http://localhost:8080/")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start server: %v\n", err.Error())
		return
	}
}
