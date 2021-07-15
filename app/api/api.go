package api

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type RequestFunction func(ctx *Ctx)
type Ctx struct {
	DB             *gorm.DB
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
