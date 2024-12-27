package main

import (
	"log"
	"net/http"
	"github.com/nikhst/rpn/internal/application"
)

func main() {
	http.HandleFunc("/api/v1/calculate", application.Answer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
