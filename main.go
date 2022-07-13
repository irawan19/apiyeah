package main

import (
	"apiyeah/router"
	"log"
	"net/http"
)

// @title API Yeah
// @version 2.0
// @description API for YEAH Apps

// @host api.yeah.biz.id
// @BasePath /api/v2
func main() {
	r := router.Router()

	log.Println(http.ListenAndServe(":8080", r))
}
