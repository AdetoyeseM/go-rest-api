package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"first-rest-api/models"

	"github.com/gorilla/mux"
)

var Books []models.Book

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Context-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(Books); 
	err != nil {
		log.Printf("ERROR: Failed to encode books in GetAllBooks: %v", err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("SUCCESS: Retrieved all books, total count: %d", len(Books))
}

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Context-Type", "application/json")
	var parameter = mux.Vars(request)

	for _, item := range Books {
		if item.ID == parameter["id"] {
			if err := json.NewEncoder(writer).Encode(item); err != nil {
				log.Printf("ERROR: Failed to encode book in GetBookByID: %v", err)
				http.Error(writer, "Internal server error", http.StatusInternalServerError)
				return
			}
			log.Printf("SUCCESS: Retrieved book with ID: %s", parameter["id"])
			return
		}
	}

	log.Printf("ERROR: Book not found with ID: %s", parameter["id"])
	http.Error(writer, "Book not found", http.StatusNotFound)

}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Context-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		log.Printf("ERROR: Failed to decode request body in CreateBook: %v", err)
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}
	Books = append(Books, book)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		log.Printf("ERROR: Failed to encode book in CreateBook: %v", err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("SUCCESS: Created new book with ID: %s", book.ID)

}

func UpdateBookByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Context-Type", "application/json")

	parameter := mux.Vars(request)

	for index, item := range Books {

		if item.ID == parameter["id"] {
			Books = append(Books[:index], Books[index+1:]...)

			var book models.Book

			if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
				log.Printf("ERROR: Failed to decode request body in UpdateBookByID: %v", err)
				http.Error(writer, "Invalid request body", http.StatusBadRequest)
				return
			}
			book.ID = parameter["id"]
			Books = append(Books, book)
			if err := json.NewEncoder(writer).Encode(book); err != nil {
				log.Printf("ERROR: Failed to encode book in UpdateBookByID: %v", err)
				http.Error(writer, "Internal server error", http.StatusInternalServerError)
				return
			}
			log.Printf("SUCCESS: Updated book with ID: %s", parameter["id"])
			return
		}

		log.Printf("ERROR: Book not found with ID: %s", parameter["id"])
		http.Error(writer, "Book Not Found", http.StatusNotFound)

	}

}

func DeleteBookByID(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	found := false
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			found = true
			log.Printf("SUCCESS: Deleted book with ID: %s", params["id"])
			break
		}
	}
	if !found {
		log.Printf("ERROR: Book not found for deletion with ID: %s", params["id"])
	}
	if err := json.NewEncoder(writer).Encode(Books); err != nil {
		log.Printf("ERROR: Failed to encode books in DeleteBookByID: %v", err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

}
