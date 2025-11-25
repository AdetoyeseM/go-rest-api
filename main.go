package main

import (
	"first-rest-api/db"
	"first-rest-api/handlers"
	"first-rest-api/models"
	"first-rest-api/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	var a = os.Getenv("SMTP_EMAIL")
	log.Print("TESTING ENV -- ", a)
	db.Connect()

	var server = router.SetupRoutes()
	handlers.Books = append(handlers.Books, models.Book{ID: "1", Title: "This is Book one", Content: "This is the content for Book one"})
	handlers.Books = append(handlers.Books, models.Book{ID: "2", Title: "This is Book 2", Content: "This is the content for Book 2"})
	handlers.Books = append(handlers.Books, models.Book{ID: "3", Title: "This is Book 3", Content: "This is the content for Book 3"})
	handlers.Books = append(handlers.Books, models.Book{ID: "4", Title: "This is Book 4", Content: "This is the content for Book 4"})
	handlers.Books = append(handlers.Books, models.Book{ID: "5", Title: "This is Book 5", Content: "This is the content for Book 5"})
	log.Println("SUCCESS: Server starting on port :8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
