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
// @version 1.0
// @description API Using Golang

// @host api.yeah.biz.id
// @BasePath /api/v1
func main() {
	r := router.Router()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server dijalankan pada port 9100...")

	log.Fatal(srv.ListenAndServe())
}
