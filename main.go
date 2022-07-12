package main

import (
	"apiyeah/router"
	"fmt"
	"log"
	"net/http"
)

// @title API Yeah
// @version 2.0
// @description API for YEAH Apps

// @host localhost:8080
// @BasePath /api/v2
func main() {
	r := router.Router()
	// // fs := http.FileServer(http.Dir("build"))
	// // http.Handle("/", fs)

	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
