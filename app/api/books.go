package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohslyfox/go-restful-example/app/db"
)

func GetAllBooks(ctx *Ctx) {
	books := []db.Book{}
	ctx.Database.Find(&books)
	respondJSON(ctx.ResponseWriter, http.StatusOK, books)
}

func InsertOrUpdateBook(ctx *Ctx) {
	book := db.Book{}

	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&book); err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer ctx.Request.Body.Close()

	if err := ctx.Database.Save(&book).Error; err != nil {
		respondError(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusCreated, &book)
}

func DeleteBook(ctx *Ctx) {
	vars := mux.Vars(ctx.Request)

	title := vars["title"]
	book := GetBook(ctx, title)
	if book == nil {
		return
	}

	if err := ctx.Database.Delete(&book).Error; err != nil {
		respondError(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusNoContent, nil)
}

func GetBook(ctx *Ctx, title string) *db.Book {
	res := db.Book{}
	if err := ctx.Database.First(&res, db.Book{Title: title}).Error; err != nil {
		respondError(ctx.ResponseWriter, http.StatusNotFound, err.Error())
		return nil
	}
	return &res
}
