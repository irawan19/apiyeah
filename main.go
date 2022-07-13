package main

import (
	"apiyeah/controller"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API Yeah
// @version 2.0
// @description API for YEAH Apps

// @host api.yeah.biz.id
// @BasePath /api/v2
func main() {
	router := mux.NewRouter()

	//Data
	router.HandleFunc("/api/v2/data/konfigurasiaplikasi", controller.AmbilKonfigurasiAplikasi).Methods("GET")
	router.HandleFunc("/api/v2/data/jeniskelamin", controller.AmbilJenisKelamin).Methods("GET")
	router.HandleFunc("/api/v2/data/pembayaran", controller.AmbilPembayaran).Methods("GET")
	router.HandleFunc("/api/v2/data/statuspembayaran", controller.AmbilPembayaran).Methods("GET")

	//Event
	router.HandleFunc("/api/v2/event", controller.AmbilSemuaEvent).Methods("GET")
	router.HandleFunc("/api/v2/event/detail/{id}", controller.AmbilDetailEvent).Methods("GET")
	router.HandleFunc("/api/v2/event/cekticket/{bookingcode}", controller.CekTicket).Methods("POST")
	router.HandleFunc("/api/v2/event/registrasi", controller.Registrasi).Methods("POST")
	router.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
