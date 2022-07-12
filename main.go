package main

import (
	"apiyeah/router"
	"fmt"
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

	http.ListenAndServe(":8080", nil)
}
