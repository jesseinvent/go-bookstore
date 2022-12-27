package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/jesseinvent/go-bookstore/pkg/utils"
	"github.com/jesseinvent/go-bookstore/pkg/models"
)

var NewBook models.Book;

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks();

	res, err := json.Marshal(newBooks);

	if err != nil {
		fmt.Println(err);
	}

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);

	w.Write(res); 
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);

	bookId := vars["id"];

	id, err := strconv.ParseInt(bookId, 0, 0);

	if err != nil {
		fmt.Println(err);
	}

	book, _ := models.GetBookById(id);

	res, _ := json.Marshal(book);

	w.Header().Set("Content-Type", "application/json");

	w.WriteHeader(http.StatusOK);

	w.Write(res);
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{};

	utils.ParseBody(r, CreateBook);

	book := CreateBook.CreateBook();

	res, _ := json.Marshal(book);

	w.Header().Set("Content-Type", "application/json");

	w.WriteHeader(http.StatusOK);

	w.Write(res);
}

func DeleteBook(w http.ResponseWriter, r *http.Request) { 
	vars := mux.Vars(r);
	
	bookId := vars["id"];

	id, err := strconv.ParseInt(bookId, 0, 0);

	if err != nil {
		fmt.Println(err)
	}

	book := models.DeleteBook(id);

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json");

	w.Write(res);

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{};

	utils.ParseBody(r, updateBook);

	vars := mux.Vars(r);

	bookId := vars["id"];

	id, err := strconv.ParseInt(bookId, 0, 0);

	if err != nil {
		fmt.Println(err);
	}

	// Get Book
	book, db := models.GetBookById(id);

	if updateBook.Name != "" {
		book.Name = updateBook.Name	
	}
 
	if updateBook.Author != "" {
		book.Author = updateBook.Author	
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication	
	}

	db.Save(&book);

	res, _ := json.Marshal(book);

	w.Header().Set("Content-Type", "application/json");

	w.WriteHeader(http.StatusOK);

	w.Write(res);

}