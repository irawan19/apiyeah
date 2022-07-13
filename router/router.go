package router

import (
	"apiyeah/controller"
	"net/http"

	"github.com/gorilla/mux"

	_ "apiyeah/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() *mux.Router {

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
	http.Handle("/", router)

	return router
}
