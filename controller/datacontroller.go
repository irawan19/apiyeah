package controller

import (
	"encoding/json"

	"log"
	"net/http"

	"apiyeah/models"

	_ "github.com/lib/pq"
)

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

type ResponseSosialMedia struct {
	Status  string `json:"status"`
	Message []models.SosialMedia
}

type ResponseMetaTag struct {
	Status  string `json:"status"`
	Message []models.MetaTag
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

// @Summary ambil sosial media
// @Schemes
// @Description ambil sosial media
// @Tags Data
// @Accept json
// @Produce json
// @Success 200 {string} AmbilSosialMedia
// @Router /data/sosialmedia [get]
func AmbilSosialMedia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	SosialMedia, err := models.AmbilSemuaSosialMedia()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponseSosialMedia
	response.Status = "Sukses"
	response.Message = SosialMedia

	json.NewEncoder(w).Encode(response)
}

// @Summary ambil meta tag
// @Schemes
// @Description ambil meta tag
// @Tags Data
// @Accept json
// @Produce json
// @Success 200 {string} AmbilMetaTag
// @Router /data/metatag [get]
func AmbilMetaTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	MetaTag, err := models.AmbilSemuaMetaTag()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response ResponseMetaTag
	response.Status = "Sukses"
	response.Message = MetaTag

	json.NewEncoder(w).Encode(response)
}
