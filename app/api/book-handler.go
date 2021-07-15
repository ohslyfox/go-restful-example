package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ohslyfox/go-restful-example/app/db"
)

func List(ctx *Ctx) {
	books := []db.Book{}
	ctx.DB.Find(&books)
	respondJSON(ctx.ResponseWriter, http.StatusOK, books)
}

func ListOne(ctx *Ctx) {
	vars := mux.Vars(ctx.Request)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}

	book, err := getBookByID(ctx, id)
	if book == nil {
		respondError(ctx.ResponseWriter, http.StatusNotFound, "Book with given id does not exist")
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusOK, book)
}

func Insert(ctx *Ctx) {
	book := db.Book{}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&book)

	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}

	if checkRatingError(ctx, &book) {
		return
	}

	if err := ctx.DB.Create(&book).Error; err != nil {
		handleSaveError(ctx, err);
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusCreated, &book)
}

func Update(ctx *Ctx) {
	vars := mux.Vars(ctx.Request)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}

	book, err := getBookByID(ctx, id)
	if book == nil {
		respondError(ctx.ResponseWriter, http.StatusNotFound, "Book with given id does not exist")
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err = decoder.Decode(&book)

	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}

	if checkRatingError(ctx, book) {
		return
	}

	book.ID = id // persist original id
	if err := ctx.DB.Save(&book).Error; err != nil {
		handleSaveError(ctx, err);
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusOK, book)
}

func Delete(ctx *Ctx) {
	vars := mux.Vars(ctx.Request)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}

	book, err := getBookByID(ctx, id)
	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusNotFound, err.Error())
		return
	}

	if err := ctx.DB.Delete(&book).Error; err != nil {
		respondError(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusNoContent, nil)
}

func DeleteAll(ctx *Ctx) {
	if err := ctx.DB.Where("1=1").Delete(&db.Book{}).Error; err != nil {
		respondError(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusNoContent, nil)
}

func CheckOut(ctx *Ctx) {
	checkInOut(ctx, false)
}

func CheckIn(ctx *Ctx) {
	checkInOut(ctx, true)
}

func checkInOut(ctx *Ctx, checkedIn bool) {
	vars := mux.Vars(ctx.Request)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, err.Error())
		return
	}

	book, err := getBookByID(ctx, id)
	if err != nil {
		respondError(ctx.ResponseWriter, http.StatusNotFound, err.Error())
		return
	}
	book.Status = checkedIn
	if err := ctx.DB.Save(&book).Error; err != nil {
		handleSaveError(ctx, err);
		return
	}
	respondJSON(ctx.ResponseWriter, http.StatusOK, book)
}

func checkRatingError(ctx *Ctx, book *db.Book) bool {
	if book != nil && (book.Rating < 1 || book.Rating > 3) {
		respondError(ctx.ResponseWriter, http.StatusBadRequest, "Book rating must be between 1-3")
		return true
	}
	return false
}

func handleSaveError(ctx *Ctx, err error) {
	if strings.Contains(err.Error(), "UNIQUE") {
		respondError(ctx.ResponseWriter, http.StatusConflict, err.Error())
		return
	}
	respondError(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
}

func getBookByID(ctx *Ctx, id uint64) (*db.Book, error) {
	return getBook(ctx, db.Book{ID: id})
}

func getBook(ctx *Ctx, book db.Book) (*db.Book, error) {
	res := db.Book{}
	if err := ctx.DB.First(&res, book).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
