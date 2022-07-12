package models

import (
	"apiyeah/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type KonfigurasiAplikasi struct {
	Nama_konfigurasi_aplikasis      string `json:"nama_konfigurasi_aplikasis"`
	Keywords_konfigurasi_aplikasis  string `json:"keywords_konfigurasi_aplikasis"`
	Deskripsi_konfigurasi_aplikasis string `json:"deskripsi_konfigurasi_aplikasis"`
	Icon_konfigurasi_aplikasis      string `json:"icon_konfigurasi_aplikasis"`
	Logo_konfigurasi_aplikasis      string `json:"logo_konfigurasi_aplikasis"`
	Logo_text_konfigurasi_aplikasis string `json:"logo_text_konfigurasi_aplikasis"`
	Whatsapp_konfigurasi_aplikasis  string `json:"whatsapp_konfigurasi_aplikasis"`
}

type JenisKelamin struct {
	Id_jenis_kelamins   int64  `json:"id_jenis_kelamins"`
	Nama_jenis_kelamins string `json:"nama_jenis_kelamins"`
}

type Pembayaran struct {
	Id_pembayarans   int64  `json:"id_pembayarans"`
	Nama_pembayarans string `json:"nama_pembayarans"`
}

type StatusPembayaran struct {
	Id_status_pembayarans   int64  `json:"id_status_pembayarans"`
	Nama_status_pembayarans string `json:"nama_status_pembayarans"`
}

type Event struct {
	Id_events                 int64           `json:"id_events"`
	Tanggal_events            string          `json:"tanggal_events"`
	Gambar_events             string          `json:"gambar_events"`
	Nama_events               string          `json:"nama_events"`
	Deskripsi_events          string          `json:"deskripsi_events"`
	Disclaimer_events         string          `json:"disclaimer_events"`
	Lokasi_events             string          `json:"lokasi_events"`
	Mulai_registrasi_events   string          `json:"mulai_registrasi_events"`
	Selesai_registrasi_events string          `json:"selesai_registrasi_events"`
	Tickets                   json.RawMessage `json:"tickets_data"`
}

type CekTicket struct {
	Id_events                     int64           `json:"id_events"`
	Tanggal_events                string          `json:"tanggal_events"`
	Gambar_events                 string          `json:"gambar_events"`
	Nama_events                   string          `json:"nama_events"`
	Deskripsi_events              string          `json:"deskripsi_events"`
	Disclaimer_events             string          `json:"disclaimer_events"`
	Lokasi_events                 string          `json:"lokasi_events"`
	Id_tickets                    int64           `json:"id_tickets"`
	Nama_tickets                  string          `json:"nama_tickets"`
	Deskripsi_tickets             string          `json:"deskripsi_tickets"`
	Keterangan_tickets            string          `json:"keterangan_tickets"`
	Tanggal_registrasi_events     string          `json:"tanggal_registrasi_events"`
	Jumlah_registrasi_events      string          `json:"jumlah_registrasi_events"`
	Harga_regisrasi_events        string          `json:"harga_registrasi_events"`
	Total_harga_registrasi_events string          `json:"total_harga_pembayarans"`
	Id_pembayarans                string          `json:"id_pembayarans"`
	Nama_pembayarans              string          `json:"nama_pembayarans"`
	Id_status_pembayarans         string          `json:"id_status_pembayarans"`
	Nama_status_pembayarans       string          `json:"nama_status_pembayarans"`
	Registrasi_data               json.RawMessage `json:"registrasi_data"`
}

type EventDetail struct {
	Id_events                 int64           `json:"id_events"`
	Tanggal_events            string          `json:"tanggal_events"`
	Gambar_events             string          `json:"gambar_events"`
	Nama_events               string          `json:"nama_events"`
	Deskripsi_events          string          `json:"deskripsi_events"`
	Disclaimer_events         string          `json:"disclaimer_events"`
	Lokasi_events             string          `json:"lokasi_events"`
	Mulai_registrasi_events   string          `json:"mulai_registrasi_events"`
	Selesai_registrasi_events string          `json:"selesai_registrasi_events"`
	Tickets                   json.RawMessage `json:"tickets_data"`
}

type RegistrasiEventJson struct {
	Tickets_id                         int64                   `json:"tickets_id"`
	Pembayarans_id                     int64                   `json:"pembayarans_id"`
	Status_pembayarans_id              int64                   `json:"status_pembayarans_id"`
	Jumlah_registrasi_events           int64                   `json:"jumlah_registrasi_events"`
	Total_harga_registrasi_events      int64                   `json:"total_harga_registrasi_events"`
	Bukti_pembayaran_registrasi_events string                  `json:"bukti_pembayaran_registrasi_events"`
	No_registrasi_events               string                  `json:"no_registrasi_events"`
	Harga_registrasi_events            int64                   `json:"harga_registrasi_events"`
	Registrasi                         []RegistrasiEventDetail `json:"registrasi"`
}

type RegistrasiEvent struct {
	Tickets_id                         int64  `json:"tickets_id"`
	Pembayarans_id                     int64  `json:"pembayarans_id"`
	Status_pembayarans_id              int64  `json:"status_pembayarans_id"`
	Jumlah_registrasi_events           int64  `json:"jumlah_registrasi_events"`
	Total_harga_registrasi_events      int64  `json:"total_harga_registrasi_events"`
	Bukti_pembayaran_registrasi_events string `json:"bukti_pembayaran_registrasi_events"`
	No_registrasi_events               string `json:"no_registrasi_events"`
	Harga_registrasi_events            int64  `json:"harga_registrasi_events"`
}

type RegistrasiEventDetail struct {
	Jenis_kelamins_id        int64  `json:"jenis_kelamins_id"`
	Nama_registrasi          string `json:"nama_registrasi"`
	Email_registrasi         string `json:"email_registrasi"`
	Telepon_registrasi       string `json:"telepon_registrasi"`
	Tanggal_lahir_registrasi string `json:"tanggal_lahir_registrasi"`
}

func AmbilSatuKonfigurasiAplikasi() (KonfigurasiAplikasi, error) {
	db := config.CreateConnection()
	defer db.Close()

	var id = 1
	var konfigurasiAplikasi KonfigurasiAplikasi

	sqlStatement := `SELECT
							nama_konfigurasi_aplikasis,
							keywords_konfigurasi_aplikasis,
							deskripsi_konfigurasi_aplikasis,
							icon_konfigurasi_aplikasis,
							logo_konfigurasi_aplikasis,
							logo_text_konfigurasi_aplikasis,
							whatsapp_konfigurasi_aplikasis
						FROM master_konfigurasi_aplikasis
						WHERE id_konfigurasi_aplikasis = $1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&konfigurasiAplikasi.Nama_konfigurasi_aplikasis,
		&konfigurasiAplikasi.Keywords_konfigurasi_aplikasis,
		&konfigurasiAplikasi.Deskripsi_konfigurasi_aplikasis,
		&konfigurasiAplikasi.Icon_konfigurasi_aplikasis,
		&konfigurasiAplikasi.Logo_konfigurasi_aplikasis,
		&konfigurasiAplikasi.Logo_text_konfigurasi_aplikasis,
		&konfigurasiAplikasi.Whatsapp_konfigurasi_aplikasis,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return konfigurasiAplikasi, nil
	case nil:
		return konfigurasiAplikasi, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return konfigurasiAplikasi, err
}

func AmbilSemuaJenisKelamin() ([]JenisKelamin, error) {
	db := config.CreateConnection()

	defer db.Close()

	var jeniskelamins []JenisKelamin

	sqlStatement := `SELECT
						id_jenis_kelamins,
						nama_jenis_kelamins
					FROM master_jenis_kelamins`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var jeniskelamin JenisKelamin

		err = rows.Scan(&jeniskelamin.Id_jenis_kelamins,
			&jeniskelamin.Nama_jenis_kelamins,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		jeniskelamins = append(jeniskelamins, jeniskelamin)
	}

	return jeniskelamins, err
}

func AmbilSemuaPembayaran() ([]Pembayaran, error) {
	db := config.CreateConnection()

	defer db.Close()

	var pembayarans []Pembayaran

	sqlStatement := `SELECT
						id_pembayarans,
						nama_pembayarans
					FROM master_pembayarans`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var pembayaran Pembayaran

		err = rows.Scan(&pembayaran.Id_pembayarans,
			&pembayaran.Nama_pembayarans,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		pembayarans = append(pembayarans, pembayaran)
	}

	return pembayarans, err
}

func AmbilSemuaStatusPembayaran() ([]StatusPembayaran, error) {
	db := config.CreateConnection()

	defer db.Close()

	var statuspembayarans []StatusPembayaran

	sqlStatement := `SELECT
						id_status_pembayarans,
						nama_status_pembayarans
					FROM master_status_pembayarans`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var statuspembayaran StatusPembayaran

		err = rows.Scan(&statuspembayaran.Id_status_pembayarans,
			&statuspembayaran.Nama_status_pembayarans,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		statuspembayarans = append(statuspembayarans, statuspembayaran)
	}

	return statuspembayarans, err
}

func AmbilSemuaEvent() ([]Event, error) {
	db := config.CreateConnection()

	defer db.Close()

	var events []Event

	sqlStatement := `SELECT
						id_events,
						tanggal_events,
						gambar_events,
						nama_events,
						deskripsi_events,
						disclaimer_events,
						lokasi_events,
						mulai_registrasi_events,
						selesai_registrasi_events,
						json_agg(t.*) AS tickets_data
					FROM master_events e
					JOIN master_tickets t ON t.events_id=e.id_events
					WHERE e.status_hapus_events=false
					AND e.tanggal_events > current_date
					GROUP BY e.id_events
					ORDER BY e.tanggal_events desc`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var event Event

		err = rows.Scan(&event.Id_events,
			&event.Tanggal_events,
			&event.Gambar_events,
			&event.Nama_events,
			&event.Deskripsi_events,
			&event.Disclaimer_events,
			&event.Lokasi_events,
			&event.Mulai_registrasi_events,
			&event.Selesai_registrasi_events,
			&event.Tickets,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		events = append(events, event)
	}

	return events, err
}

func AmbilSatuEvent(id int64) (EventDetail, error) {
	db := config.CreateConnection()
	defer db.Close()

	var eventDetail EventDetail

	sqlStatement := `SELECT
							id_events,
							tanggal_events,
							gambar_events,
							nama_events,
							deskripsi_events,
							disclaimer_events,
							lokasi_events,
							mulai_registrasi_events,
							selesai_registrasi_events,
							json_agg(t.*) as tickets_data
						FROM master_events e
						JOIN master_tickets t ON t.events_id=e.id_events
						WHERE e.status_hapus_events=false
						AND e.id_events=$1
						GROUP BY id_events`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&eventDetail.Id_events,
		&eventDetail.Tanggal_events,
		&eventDetail.Gambar_events,
		&eventDetail.Nama_events,
		&eventDetail.Deskripsi_events,
		&eventDetail.Disclaimer_events,
		&eventDetail.Lokasi_events,
		&eventDetail.Mulai_registrasi_events,
		&eventDetail.Selesai_registrasi_events,
		&eventDetail.Tickets,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return eventDetail, nil
	case nil:
		return eventDetail, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return eventDetail, err
}

func AmbilSatuTicket(booking_code string) (CekTicket, error) {
	db := config.CreateConnection()
	defer db.Close()

	var cekTicket CekTicket

	sqlStatement := `SELECT
							id_events,
							tanggal_events,
							gambar_events,
							nama_events,
							deskripsi_events,
							disclaimer_events,
							lokasi_events,
							id_tickets,
							nama_tickets,
							deskripsi_tickets,
							keterangan_tickets,
							re.created_at AS tanggal_registrasi_events,
							jumlah_registrasi_events,
							harga_registrasi_events,
							total_harga_registrasi_events,
							id_pembayarans,
							nama_pembayarans,
							id_status_pembayarans,
							nama_status_pembayarans,
							json_agg(red.*) as registrasi_data
						FROM registrasi_events re
						JOIN registrasi_event_details red ON red.registrasi_events_id=re.id_registrasi_events
						JOIN master_tickets t ON t.id_tickets=re.tickets_id
						JOIN master_events e ON e.id_events=t.events_id
						JOIN master_pembayarans p ON p.id_pembayarans=re.pembayarans_id
						JOIN master_status_pembayarans sp ON sp.id_status_pembayarans=re.status_pembayarans_id
						JOIN master_jenis_kelamins jk ON jk.id_jenis_kelamins=red.jenis_kelamins_id
						WHERE re.no_registrasi_events=$1
						GROUP BY re.id_registrasi_events, e.id_events, t.id_tickets, p.id_pembayarans, sp.id_status_pembayarans, jk.id_jenis_kelamins`

	row := db.QueryRow(sqlStatement, booking_code)

	err := row.Scan(&cekTicket.Id_events,
		&cekTicket.Tanggal_events,
		&cekTicket.Gambar_events,
		&cekTicket.Nama_events,
		&cekTicket.Deskripsi_events,
		&cekTicket.Disclaimer_events,
		&cekTicket.Lokasi_events,
		&cekTicket.Id_tickets,
		&cekTicket.Nama_tickets,
		&cekTicket.Deskripsi_tickets,
		&cekTicket.Keterangan_tickets,
		&cekTicket.Tanggal_registrasi_events,
		&cekTicket.Jumlah_registrasi_events,
		&cekTicket.Harga_regisrasi_events,
		&cekTicket.Total_harga_registrasi_events,
		&cekTicket.Id_pembayarans,
		&cekTicket.Nama_pembayarans,
		&cekTicket.Id_status_pembayarans,
		&cekTicket.Nama_status_pembayarans,
		&cekTicket.Registrasi_data,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return cekTicket, nil
	case nil:
		return cekTicket, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return cekTicket, err
}

func TambahDataRegistrasiEvent(registrasievent RegistrasiEvent) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatementRegistrasiEvent := `INSERT INTO registrasi_events (
										tickets_id,
										pembayarans_id,
										status_pembayarans_id,
										jumlah_registrasi_events,
										total_harga_registrasi_events,
										bukti_pembayaran_registrasi_events,
										no_registrasi_events,
										harga_registrasi_events,
										created_at,
										updated_at
									)
									VALUES ($1,$2,$3,$4,$5,$6,$7,$8, NOW(), NOW()) RETURNING id_registrasi_events`

	var idregistrasievent int64
	errRegistrasiEvent := db.QueryRow(sqlStatementRegistrasiEvent,
		registrasievent.Tickets_id,
		registrasievent.Pembayarans_id,
		registrasievent.Status_pembayarans_id,
		registrasievent.Jumlah_registrasi_events,
		registrasievent.Total_harga_registrasi_events,
		registrasievent.Bukti_pembayaran_registrasi_events,
		registrasievent.No_registrasi_events,
		registrasievent.Harga_registrasi_events,
	).Scan(&idregistrasievent)

	if errRegistrasiEvent != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", errRegistrasiEvent)
	}
	fmt.Printf("Insert data single record %v", idregistrasievent)

	return idregistrasievent
}

func TambahDataRegistrasiEventDetail(id int64, registrasieventdetail RegistrasiEventDetail) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatementRegistrasiEventDetail := `INSERT INTO registrasi_event_details (
										registrasi_events_id,
										jenis_kelamins_id,
										nama_registrasi_event_details,
										email_registrasi_event_details,
										telepon_registrasi_event_details,
										tanggal_lahir_registrasi_event_details,
										created_at,
										updated_at
									)
									VALUES ($1,$2,$3,$4,$5,$6, NOW(), NOW()) RETURNING id_registrasi_event_details`

	var idregistrasieventdetail int64

	_, errRegistrasiEvent := db.Exec(sqlStatementRegistrasiEventDetail,
		id,
		registrasieventdetail.Jenis_kelamins_id,
		registrasieventdetail.Nama_registrasi,
		registrasieventdetail.Email_registrasi,
		registrasieventdetail.Telepon_registrasi,
		registrasieventdetail.Tanggal_lahir_registrasi,
	)

	if errRegistrasiEvent != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", errRegistrasiEvent)
	}
	fmt.Printf("Insert data single record %v", idregistrasieventdetail)

	return idregistrasieventdetail
}