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

// @Summary ambil konfigurasi aplikasi
// @Schemes
// @Description ambil konfigurasi aplikasi
// @Tags Data
// @Accept json
// @Produce json
// @Success 200 {string} AmbilKonfigurasiAplikasi
// @Router /data/konfigurasiaplikasi [get]
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

// @Summary ambil jenis kelamin
// @Schemes
// @Description ambil jenis kelamin
// @Tags Data
// @Accept json
// @Produce json
// @Success 200 {string} AmbilJenisKelamin
// @Router /data/jeniskelamin [get]
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

// @Summary ambil pembayaran
// @Schemes
// @Description ambil pembayaran
// @Tags Data
// @Accept json
// @Produce json
// @Success 200 {string} AmbilPembayaran
// @Router /data/pembayaran [get]
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

// @Summary ambil status pembayaran
// @Schemes
// @Description ambil status pembayaran
// @Tags Data
// @Accept json
// @Produce json
// @Success 200 {string} AmbilStatusPembayaran
// @Router /data/statuspembayaran [get]
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

// @Summary ambil semua event
// @Schemes
// @Description ambil semua event
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {string} AmbilSemuaEvent
// @Router /event [get]
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

// @Summary ambil detail event
// @Schemes
// @Description ambil detail event
// @Tags Event
// @Accept json
// @Produce json
// @Param id query string true "id event"
// @Success 200 {string} AmbilDetailEvent
// @Router /event/detail/{id} [get]
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

// @Summary cek ticket
// @Schemes
// @Description cek ticket
// @Tags Event
// @Accept json
// @Produce json
// @Param booking_code query string true "booking code"
// @Success 200 {string} AmbilDetailEvent
// @Router /event/cekticket/ [post]
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

// @Summary registrasi
// @Schemes
// @Description registrasi
// @Tags Event
// @Accept json
// @Produce json
// @Param registrasi body object true "registrasi"
// @Success 200 {string} AmbilDetailEvent
// @Router /event/registrasi [post]
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

	ticketdata := models.KalkulasiTicketData{
		Jumlah_registrasi: regjson.Jumlah_registrasi_events,
	}
	models.KalkulasiKuotaTicket(regjson.Tickets_id, ticketdata)

	res := response{
		Status:  "Sukses",
		Message: "Data registrasi telah ditambahkan ",
	}
	json.NewEncoder(w).Encode(res)
}
