package main

import (
	"chap3/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
