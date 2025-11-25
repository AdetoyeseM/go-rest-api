package router

import (
	"first-rest-api/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	var server = mux.NewRouter()

	server.HandleFunc("/getAllBooks", handlers.AuthMiddleware(handlers.GetAllBooks)).Methods("GET")
	server.HandleFunc("/getBookByID/{id}", handlers.AuthMiddleware(handlers.GetBookByID)).Methods("GET")
	server.HandleFunc("/createBook", handlers.AuthMiddleware(handlers.CreateBook)).Methods("POST")
	server.HandleFunc("/updateBookByID/{id}", handlers.AuthMiddleware(handlers.UpdateBookByID)).Methods("PUT")

	server.HandleFunc("/deleteBook/{id}", handlers.AuthMiddleware(handlers.DeleteBookByID)).Methods("DELETE")
	server.HandleFunc("/register", handlers.Register).Methods("POST")
	server.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	return server
}
