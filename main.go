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

// @host api.yeah.biz.id
// @BasePath /api/v2
func main() {
	r := router.Router()

	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
