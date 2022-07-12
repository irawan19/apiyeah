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
	router.HandleFunc("/api/v2/data/konfigurasiaplikasi", controller.AmbilKonfigurasiAplikasi).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v2/data/jeniskelamin", controller.AmbilJenisKelamin).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v2/data/pembayaran", controller.AmbilPembayaran).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v2/data/statuspembayaran", controller.AmbilPembayaran).Methods("GET", "OPTIONS")

	//Event
	router.HandleFunc("/api/v2/event", controller.AmbilSemuaEvent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v2/event/detail/{id}", controller.AmbilDetailEvent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v2/event/cekticket/{bookingcode}", controller.CekTicket).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v2/event/registrasi", controller.Registrasi).Methods("POST", "OPTIONS")
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	return router
}
