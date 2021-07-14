package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohslyfox/go-restful-example/app/api"
	"github.com/ohslyfox/go-restful-example/app/db"
	"gorm.io/gorm"
)

var GET string = "GET"
var POST string = "POST"
var PUT string = "PUT"
var DELETE string = "DELETE"

type App struct {
	Router   *mux.Router
	Database *gorm.DB
}

var Instance *App

func GetInstance() *App {
	if Instance == nil {
		Instance = InitializeApp()
	}
	return Instance
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":3000", a.Router))
}

func InitializeApp() *App {
	res := &App{}
	res.Database = db.GetDatabaseInstance()
	res.Router = mux.NewRouter()
	res.AddRoute("/books", GET, api.GetAllBooks)
	res.AddRoute("/books", POST, api.InsertOrUpdateBook)
	res.AddRoute("/books", DELETE, api.DeleteBook)
	return res
}

func (app *App) AddRoute(path string, reqType string, fn api.RequestFunction) {
	handlerFunc := GetHandlerFunctionFromApiFunction(fn, app.Database)
	app.Router.HandleFunc(path, handlerFunc).Methods(reqType)
}

func GetHandlerFunctionFromApiFunction(fn api.RequestFunction, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(
			&api.Ctx{
				Database:       db,
				Request:        r,
				ResponseWriter: w,
			},
		)
	}
}
