package router

import (
	"apiyeah/controller"

	"github.com/gorilla/mux"

	_ "apiyeah/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//Data
	router.HandleFunc("/api/v1/data/konfigurasiaplikasi", controller.AmbilKonfigurasiAplikasi).Methods("GET")
	router.HandleFunc("/api/v1/data/jeniskelamin", controller.AmbilJenisKelamin).Methods("GET")
	router.HandleFunc("/api/v1/data/pembayaran", controller.AmbilPembayaran).Methods("GET")
	router.HandleFunc("/api/v1/data/statuspembayaran", controller.AmbilPembayaran).Methods("GET")
	router.HandleFunc("/api/v1/data/sosialmedia", controller.AmbilSosialMedia).Methods("GET")
	router.HandleFunc("/api/v1/data/metatag", controller.AmbilMetaTag).Methods("GET")

	//Event
	router.HandleFunc("/api/v1/event", controller.AmbilSemuaEvent).Methods("GET")
	router.HandleFunc("/api/v1/event/detail/{id}", controller.AmbilDetailEvent).Methods("GET")
	router.HandleFunc("/api/v1/event/cekticket/{bookingcode}", controller.CekTicket).Methods("POST")
	router.HandleFunc("/api/v1/event/registrasi", controller.Registrasi).Methods("POST")
	router.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	return router
}
