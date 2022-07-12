package router

import (
	"apiyeah/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//Data
	router.HandleFunc("/api/v1/data/konfigurasiaplikasi", controller.AmbilKonfigurasiAplikasi).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/data/jeniskelamin", controller.AmbilJenisKelamin).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/data/pembayaran", controller.AmbilPembayaran).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/data/statuspembayaran", controller.AmbilPembayaran).Methods("GET", "OPTIONS")

	//Event
	router.HandleFunc("/api/v1/event", controller.AmbilSemuaEvent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/event/detail/{id}", controller.AmbilDetailEvent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/event/cekticket/{bookingcode}", controller.CekTicket).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/event/registrasi", controller.Registrasi).Methods("POST", "OPTIONS")

	return router
}
