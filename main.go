package main

import (
	"apiyeah/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

const ()

// @title API Yeah
// @version 2.0
// @description API for YEAH Apps

// @host api.yeah.biz.id
// @BasePath /api/v2
func main() {
	r := router.Router()
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server dijalankan pada port 9100...")

	log.Fatal(srv.ListenAndServe())
}
