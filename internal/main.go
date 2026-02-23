package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"proyectoapi/internal/service"
	"proyectoapi/internal/store"
	"proyectoapi/internal/transport"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := "CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, author TEXT NOT NULL)"

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBook)

	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	fmt.Println("API Endpoints:")
	fmt.Println("GET /books - Obtiene todos los libros")
	fmt.Println("POST /books - Crea un nuevo libro")
	fmt.Println("GET /books/{id} - Obtiene un libro por su ID")
	fmt.Println("PUT /books/{id} - Actualiza un libro por su ID")
	fmt.Println("DELETE /books/{id} - Elimina un libro por su ID")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
