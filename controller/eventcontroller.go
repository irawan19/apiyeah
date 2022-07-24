package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"log"
	"net/http"

	"apiyeah/helper"
	"apiyeah/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type responseRegistrasi struct {
	Status      string `json:"status,omitempty"`
	Message     string `json:"message,omitempty"`
	BookingCode string `json:"booking_code,omitempty"`
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
// @Param id path string true "id event"
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
// @Router /event/cekticket [post]
func CekTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bookingcode := r.FormValue("booking_code")
	cekTickets, err := models.AmbilSatuTicket(bookingcode)

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data ticket. %v", err)
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
// @Success 200 {string} RegistrasiEventDetail
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

	generateNoRegistrasi := helper.RandStringBytes(6)

	var PembayaransId int64
	var StatusPembayaransId int64
	var BuktiPembayaranRegistrasiEvents string
	if regjson.Harga_registrasi_events == 0 {
		PembayaransId = 1
		StatusPembayaransId = 2
		BuktiPembayaranRegistrasiEvents = generateNoRegistrasi
	} else {
		PembayaransId = 1
		StatusPembayaransId = 1
		BuktiPembayaranRegistrasiEvents = ""
	}

	regdata := models.RegistrasiEvent{
		Tickets_id:                         regjson.Tickets_id,
		Pembayarans_id:                     PembayaransId,
		Status_pembayarans_id:              StatusPembayaransId,
		Jumlah_registrasi_events:           regjson.Jumlah_registrasi_events,
		Bukti_pembayaran_registrasi_events: BuktiPembayaranRegistrasiEvents,
		No_registrasi_events:               generateNoRegistrasi,
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

	res := responseRegistrasi{
		Status:      "Sukses",
		Message:     "Data registrasi telah ditambahkan ",
		BookingCode: generateNoRegistrasi,
	}
	json.NewEncoder(w).Encode(res)
}

// @Summary pembayaran registrasi
// @Schemes
// @Description pembayaran registrasi
// @Tags Event
// @Accept json
// @Produce json
// @Param booking_code query string true "booking code"
// @Param id_pembayarans query string true "id pembayaran"
// @Param id_status_pembayarans query string true "id status pembayaran"
// @Success 200 {string} Pembayaran
// @Router /event/registrasi/pembayaran [post]
func Pembayaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bookingcode := r.FormValue("booking_code")
	idpembayaran, errpembayaran := strconv.Atoi(r.FormValue("id_pembayarans"))
	if errpembayaran != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", errpembayaran)
	}
	idstatuspembayaran, errstatuspembayaran := strconv.Atoi(r.FormValue("id_status_pembayarans"))
	if errstatuspembayaran != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", errstatuspembayaran)
	}

	models.UpdatePembayaran(string(bookingcode), int64(idpembayaran), int64(idstatuspembayaran))

	res := response{
		Status:  "Sukses",
		Message: "Data pembayaran registrasi telah diperbarui",
	}
	json.NewEncoder(w).Encode(res)
}

// @Summary bukti pembayaran registrasi
// @Schemes
// @Description bukti pembayaran registrasi
// @Tags Event
// @Accept json
// @Produce json
// @Param booking_code query string true "booking code"
// @Param bukti_pembayaran_registrasi_events formData file true "bukti pembayaran"
// @Success 200 {string} BuktiPembayaran
// @Router /event/registrasi/buktipembayaran [post]
func BuktiPembayaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("bukti_pembayaran_registrasi_events")
	if err != nil {
		fmt.Println("Error mendapatkan file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	bookingcode := r.FormValue("booking_code")

	buktiPembayaranLama, err := models.CekBuktiPembayaranLama(bookingcode)
	if err != nil {
		log.Fatalf("Tidak bisa mengambil data bukti pembayaran. %v", err)
	}
	if _, errexists := os.Stat(buktiPembayaranLama); err != nil {
		fmt.Println(errexists)
	} else {
		pathLama := "/var/www/html/yeah/" + buktiPembayaranLama
		errLama := os.Remove(pathLama)
		if errLama != nil {
			fmt.Println(errLama)
		}
	}

	tempFile, err := ioutil.TempFile("/var/www/html/yeah/public/uploads/bukti_pembayaran/", "apiupload-*.png")
	tempFile.Chmod(0777)
	buktipembayaran := strings.ReplaceAll(tempFile.Name(), "/var/www/html/yeah/", "")
	models.UpdateBuktiPembayaran(string(bookingcode), string(buktipembayaran))

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	res := response{
		Status:  "Sukses",
		Message: "Bukti pembayaran berhasil diupload",
	}
	json.NewEncoder(w).Encode(res)
}

// @Summary verifikasi kedatangan
// @Schemes
// @Description verifikasi kedatangan
// @Tags Event
// @Accept json
// @Produce json
// @Param booking_code query string true "booking code"
// @Param kode_scanner_events query string true "kode scanner event"
// @Success 200 {string} VerifikasiKedatangan
// @Router /event/verifikasikedatangan [post]
func VerifikasiKedatangan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bookingcode := r.FormValue("booking_code")
	kodescanner := r.FormValue("kode_scanner_events")

	cekkodescanner, err := models.CekKodeScanner(bookingcode, kodescanner)
	if err != nil {
		fmt.Println("data: ", err)
	} else {
		if cekkodescanner != 0 {
			models.UpdateKedatangan(string(bookingcode))
			res := response{
				Status:  "Error",
				Message: "Kedatangan telah diverifikasi",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			res := response{
				Status:  "Error",
				Message: "Data kode events salah!",
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}
