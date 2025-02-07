package router

import (
	handlers "bookapi/pkg/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	return router
}

func StartServer() {
	router := InitializeRoutes()
	http.Handle("/", router)
	fmt.Println("Book api server started.......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
