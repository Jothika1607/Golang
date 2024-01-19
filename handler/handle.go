package main

import (
	"fmt"
	"net/http"
)

type CustomHandler struct{}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is a custom handler!")
}

func main() {

	http.Handle("/custom", &CustomHandler{})
	http.ListenAndServe(":8080", nil)
}
