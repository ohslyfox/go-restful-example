package main

import (
	"github.com/ohslyfox/go-restful-example/app"
)

func main() {
	app := app.GetInstance()
	app.Run()
}
