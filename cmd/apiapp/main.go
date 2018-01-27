package main

import (
	"github.com/antonybudianto/go-starter/app"
)

func main() {
	a := app.App{}
	a.Initialize("root", "hello", "rest_api_example")

	a.Run(":8000")
}
