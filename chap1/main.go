package main

import (
	"golang/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", myapp.NewHttpHandler())
}
