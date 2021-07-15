package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohslyfox/go-restful-example/app/api"
	"github.com/ohslyfox/go-restful-example/app/db"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

var Instance *App

func GetInstance() *App {
	if Instance == nil {
		Instance = InitializeApp()
	}
	return Instance
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func InitializeApp() *App {
	res := &App{}
	res.DB = db.GetDatabaseInstance()
	res.Router = mux.NewRouter()
	res.AddRoute("/books", http.MethodGet, api.List)
	res.AddRoute("/books", http.MethodPost, api.Insert)
	res.AddRoute("/books", http.MethodDelete, api.DeleteAll)
	res.AddRoute("/books/{id:[0-9]+}", http.MethodGet, api.ListOne)
	res.AddRoute("/books/{id:[0-9]+}", http.MethodDelete, api.Delete)
	res.AddRoute("/books/{id:[0-9]+}", http.MethodPut, api.Update)
	res.AddRoute("/books/{id:[0-9]+}/checkout", http.MethodPut, api.CheckOut)
	res.AddRoute("/books/{id:[0-9]+}/checkin", http.MethodPut, api.CheckIn)
	return res
}

func (app *App) AddRoute(path string, reqType string, fn api.RequestFunction) {
	handlerFunc := GetHandlerFunctionFromApiFunction(fn, app.DB)
	app.Router.HandleFunc(path, handlerFunc).Methods(reqType)
}

func GetHandlerFunctionFromApiFunction(fn api.RequestFunction, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(
			&api.Ctx{
				DB:       db,
				Request:        r,
				ResponseWriter: w,
			},
		)
	}
}
