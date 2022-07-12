package controller

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"log"
	"net/http"

	"apiyeah/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseKonfigurasiAplikasi struct {
	Status  string                     `json:"status"`
	Message models.KonfigurasiAplikasi `json:"message"`
}

type ResponseJenisKelamin struct {
	Status  string                `json:"status"`
	Message []models.JenisKelamin `json:"message"`
}

type ResponsePembayaran struct {
	Status  string              `json:"status"`
	Message []models.Pembayaran `json:"message"`
}

type ResponseStatusPembayaran struct {
	Status  string                    `json:"status"`
	Message []models.StatusPembayaran `json:"message"`
}

type ResponseSemuaEvent struct {
	Status  string         `json:"status"`
	Message []models.Event `json:"message"`
}

type ResponseEventDetail struct {
	Status  string             `json:"status"`
	Message models.EventDetail `json:"message"`
}

type ResponseCekTicket struct {
	Status  string           `json:"status"`
	Message models.CekTicket `json:"message"`
}

func AmbilKonfigurasiAplikasi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	KonfigurasiAplikasis, err := models.AmbilSatuKonfigurasiAplikasi()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponseKonfigurasiAplikasi
	response.Status = "Sukses"
	response.Message = KonfigurasiAplikasis

	json.NewEncoder(w).Encode(response)
}

func AmbilJenisKelamin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	JenisKelamin, err := models.AmbilSemuaJenisKelamin()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponseJenisKelamin
	response.Status = "Sukses"
	response.Message = JenisKelamin

	json.NewEncoder(w).Encode(response)
}

func AmbilPembayaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	Pembayaran, err := models.AmbilSemuaPembayaran()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponsePembayaran
	response.Status = "Sukses"
	response.Message = Pembayaran

	json.NewEncoder(w).Encode(response)
}

func AmbilStatusPembayaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	StatusPembayaran, err := models.AmbilSemuaStatusPembayaran()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponseStatusPembayaran
	response.Status = "Sukses"
	response.Message = StatusPembayaran

	json.NewEncoder(w).Encode(response)
}

func AmbilSemuaEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	Events, err := models.AmbilSemuaEvent()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponseSemuaEvent
	response.Status = "Sukses"
	response.Message = Events

	json.NewEncoder(w).Encode(response)
}

func AmbilDetailEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	EventDetails, err := models.AmbilSatuEvent(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data event. %v", err)
	}

	var response ResponseEventDetail
	response.Status = "Sukses"
	response.Message = EventDetails

	json.NewEncoder(w).Encode(response)
}

func CekTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bookingcode := r.FormValue("booking_code")
	cekTickets, err := models.AmbilSatuTicket(bookingcode)

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data event. %v", err)
	}

	var response ResponseCekTicket
	response.Status = "Sukses"
	response.Message = cekTickets

	json.NewEncoder(w).Encode(response)
}

func Registrasi(w http.ResponseWriter, r *http.Request) {
	var regjson models.RegistrasiEventJson

	ambiljson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body. %v", err)
	}

	convertjson := json.Unmarshal([]byte(ambiljson), &regjson)
	if convertjson != nil {
		log.Fatalf("Tidak bisa mendecode dari request body. %v", convertjson)
	}

	regdata := models.RegistrasiEvent{
		Tickets_id:                         regjson.Tickets_id,
		Pembayarans_id:                     regjson.Pembayarans_id,
		Status_pembayarans_id:              regjson.Status_pembayarans_id,
		Jumlah_registrasi_events:           regjson.Jumlah_registrasi_events,
		Bukti_pembayaran_registrasi_events: regjson.Bukti_pembayaran_registrasi_events,
		No_registrasi_events:               regjson.No_registrasi_events,
		Harga_registrasi_events:            regjson.Harga_registrasi_events,
	}
	IdRegistrasiEvent := models.TambahDataRegistrasiEvent(regdata)

	jumlahregdetail := len(regjson.Registrasi)

	for i := 0; i < jumlahregdetail; i++ {
		regdatadetail := models.RegistrasiEventDetail{
			Jenis_kelamins_id:        regjson.Registrasi[i].Jenis_kelamins_id,
			Nama_registrasi:          regjson.Registrasi[i].Nama_registrasi,
			Email_registrasi:         regjson.Registrasi[i].Email_registrasi,
			Telepon_registrasi:       regjson.Registrasi[i].Telepon_registrasi,
			Tanggal_lahir_registrasi: regjson.Registrasi[i].Tanggal_lahir_registrasi,
		}
		models.TambahDataRegistrasiEventDetail(IdRegistrasiEvent, regdatadetail)
	}

	res := response{
		Status:  "Sukses",
		Message: "Data registrasi telah ditambahkan ",
	}
	json.NewEncoder(w).Encode(res)

}
