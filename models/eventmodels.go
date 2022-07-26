package models

import (
	"apiyeah/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

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
	Start_from_ticket_events  int64           `json:"start_form_ticket_events"`
	Tickets                   json.RawMessage `json:"tickets_data"`
}

type CekTicket struct {
	Id_events                          int64           `json:"id_events"`
	Tanggal_events                     string          `json:"tanggal_events"`
	Gambar_events                      string          `json:"gambar_events"`
	Nama_events                        string          `json:"nama_events"`
	Deskripsi_events                   string          `json:"deskripsi_events"`
	Disclaimer_events                  string          `json:"disclaimer_events"`
	Lokasi_events                      string          `json:"lokasi_events"`
	Id_tickets                         int64           `json:"id_tickets"`
	Nama_tickets                       string          `json:"nama_tickets"`
	Deskripsi_tickets                  string          `json:"deskripsi_tickets"`
	Keterangan_tickets                 string          `json:"keterangan_tickets"`
	Tanggal_registrasi_events          string          `json:"tanggal_registrasi_events"`
	Jumlah_registrasi_events           string          `json:"jumlah_registrasi_events"`
	Harga_regisrasi_events             string          `json:"harga_registrasi_events"`
	Total_harga_registrasi_events      string          `json:"total_harga_registrasi_events"`
	Bukti_pembayaran_registrasi_events string          `json:"bukti_pembayaran_registrasi_events"`
	Id_pembayarans                     string          `json:"id_pembayarans"`
	Nama_pembayarans                   string          `json:"nama_pembayarans"`
	No_rekening_pembayarans            string          `json:"no_rekening_pembayarans"`
	Nama_rekening_pembayarans          string          `json:"nama_rekening_pembayarans"`
	Logo_pembayarans                   string          `json:"logo_pembayarans"`
	Id_status_pembayarans              string          `json:"id_status_pembayarans"`
	Nama_status_pembayarans            string          `json:"nama_status_pembayarans"`
	Id_tipe_pembayarans                int64           `json:"id_tipe_pembayarans"`
	Nama_tipe_pembayarans              string          `json:"nama_tipe_pembayarans"`
	Registrasi_data                    json.RawMessage `json:"registrasi_data"`
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
	Promos                    json.RawMessage `json:"promos_data"`
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

type KalkulasiTicketData struct {
	Jumlah_registrasi int64 `json:"jumlah_registrasi"`
}

type JumlahRegistrasiData struct {
	Tickets_id            int64 `json:"tickets_id"`
	Status_pembayarans_id int64 `json:"status_pembayarans_id"`
	Jumlah_registrasi     int64 `json:"jumlah_registrasi"`
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
						(
							SELECT harga_tickets
							FROM master_tickets ts
							WHERE ts.events_id=e.id_events
							ORDER BY harga_tickets ASC
							LIMIT 1
						) AS start_from_ticket_events,
						json_agg(t.*) AS tickets_data
					FROM master_events e
					JOIN master_tickets t ON t.events_id=e.id_events
					WHERE e.status_hapus_events=false
					AND e.tanggal_events > current_date
					AND t.status_hapus_tickets=false
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
			&event.Start_from_ticket_events,
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
							json_agg(t.*) as tickets_data,
							(
								SELECT
									COALESCE(json_agg(p.*), '[]')
								FROM master_promos p
								WHERE p.events_id=e.id_events
								OR p.events_id=0
							) as promos_data
						FROM master_events e
						JOIN master_tickets t ON t.events_id=e.id_events
						WHERE e.status_hapus_events=false
						AND t.status_hapus_tickets=false
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
		&eventDetail.Promos,
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
							coalesce(bukti_pembayaran_registrasi_events, ''),
							id_pembayarans,
							nama_pembayarans,
							coalesce(no_rekening_pembayarans, ''),
							coalesce(nama_rekening_pembayarans, ''),
							coalesce(logo_pembayarans, ''),
							id_status_pembayarans,
							nama_status_pembayarans,
							id_tipe_pembayarans,
							nama_tipe_pembayarans,
							json_agg(
								json_build_object(
									'id_registrasi_event_details', red.id_registrasi_event_details,
									'nama_registrasi_event_details', red.nama_registrasi_event_details,
									'email_registrasi_event_details', red.email_registrasi_event_details,
									'telepon_registrasi_event_details', red.telepon_registrasi_event_details,
									'tanggal_lahir_registrasi_event_details', red.tanggal_lahir_registrasi_event_details,
									'id_jenis_kelamins', jk.id_jenis_kelamins,
									'nama_jenis_kelamins', jk.nama_jenis_kelamins,
									'created_at', red.created_at,
									'updated_at', red.updated_at
								)
							) AS registrasi_data
						FROM registrasi_events re
						JOIN registrasi_event_details red ON red.registrasi_events_id=re.id_registrasi_events
						JOIN master_tickets t ON t.id_tickets=re.tickets_id
						JOIN master_events e ON e.id_events=t.events_id
						JOIN master_pembayarans p ON p.id_pembayarans=re.pembayarans_id
						JOIN master_status_pembayarans sp ON sp.id_status_pembayarans=re.status_pembayarans_id
						JOIN master_tipe_pembayarans tp ON tp.id_tipe_pembayarans=p.tipe_pembayarans_id
						JOIN master_jenis_kelamins jk ON jk.id_jenis_kelamins=red.jenis_kelamins_id
						WHERE re.no_registrasi_events=$1
						GROUP BY re.id_registrasi_events, e.id_events, t.id_tickets, p.id_pembayarans, sp.id_status_pembayarans, tp.id_tipe_pembayarans`

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
		&cekTicket.Bukti_pembayaran_registrasi_events,
		&cekTicket.Id_pembayarans,
		&cekTicket.Nama_pembayarans,
		&cekTicket.No_rekening_pembayarans,
		&cekTicket.Nama_rekening_pembayarans,
		&cekTicket.Logo_pembayarans,
		&cekTicket.Id_status_pembayarans,
		&cekTicket.Nama_status_pembayarans,
		&cekTicket.Id_tipe_pembayarans,
		&cekTicket.Nama_tipe_pembayarans,
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
										status_kedatangan_registrasi_events,
										created_at,
										updated_at
									)
									VALUES ($1,$2,$3,$4,$5,$6,$7,$8, false, NOW(), NOW()) RETURNING id_registrasi_events`

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

func KalkulasiKuotaTicket(id int64, kalkulasiticket KalkulasiTicketData) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE master_tickets SET sisa_kuota_tickets=sisa_kuota_tickets-$2 WHERE id_tickets=$1`

	res, err := db.Exec(sqlStatement, id, kalkulasiticket.Jumlah_registrasi)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func UpdatePembayaran(Booking_code string, Id_pembayarans int64, Id_status_pembayarans int64) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE registrasi_events SET pembayarans_id=$2, status_pembayarans_id=$3 WHERE no_registrasi_events=$1`

	res, err := db.Exec(sqlStatement, Booking_code, Id_pembayarans, Id_status_pembayarans)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func UpdateBuktiPembayaran(Booking_code string, Bukti_pembayaran string) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE registrasi_events SET bukti_pembayaran_registrasi_events=$2 WHERE no_registrasi_events=$1`

	res, err := db.Exec(sqlStatement, Booking_code, Bukti_pembayaran)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func AmbilJumlahRegistrasi(Booking_code string) (JumlahRegistrasiData, error) {
	db := config.CreateConnection()
	defer db.Close()

	var jumlahRegistrasi JumlahRegistrasiData

	sqlStatement := `SELECT
							tickets_id,
							jumlah_registrasi_events,
							status_pembayarans_id
						FROM registrasi_events
						WHERE no_registrasi_events=$1`

	row := db.QueryRow(sqlStatement, Booking_code)

	err := row.Scan(&jumlahRegistrasi.Tickets_id,
		&jumlahRegistrasi.Jumlah_registrasi,
		&jumlahRegistrasi.Status_pembayarans_id,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return jumlahRegistrasi, nil
	case nil:
		return jumlahRegistrasi, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return jumlahRegistrasi, err
}

func CekKodeScanner(Booking_code string, Kode_scanner string) (int64, error) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `SELECT
							COUNT(kode_scanner_events) as count_kodescanner
						FROM registrasi_events re
						JOIN master_tickets t ON t.id_tickets=re.tickets_id
						JOIN master_events e ON e.id_events=t.events_id
						WHERE no_registrasi_events = $1
						AND e.kode_scanner_events=$2`
	rows, err := db.Query(sqlStatement, Booking_code, Kode_scanner)
	if err != nil {
		return 0, err
	} else {
		var count_kodescanner int64
		for rows.Next() {
			rows.Scan(&count_kodescanner)
		}
		return count_kodescanner, nil
	}
}

func CekStatusKedatangan(Booking_code string) (bool, error) {
	db := config.CreateConnection()
	defer db.Close()

	var statuskedatangandata bool

	sqlStatement := `SELECT
							status_kedatangan_registrasi_events
						FROM registrasi_events
						WHERE no_registrasi_events=$1`
	row := db.QueryRow(sqlStatement, Booking_code)
	err := row.Scan(&statuskedatangandata)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return statuskedatangandata, nil
	case nil:
		return statuskedatangandata, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return statuskedatangandata, err
}

func UpdateKedatangan(Booking_code string) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE registrasi_events SET status_kedatangan_registrasi_events=true WHERE no_registrasi_events=$1`

	res, err := db.Exec(sqlStatement, Booking_code)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func CekBuktiPembayaranLama(Booking_code string) (string, error) {
	db := config.CreateConnection()
	defer db.Close()

	var buktipembayarandata string

	sqlStatement := `SELECT
							bukti_pembayaran_registrasi_events
						FROM registrasi_events
						WHERE no_registrasi_events = $1`
	row := db.QueryRow(sqlStatement, Booking_code)
	err := row.Scan(&buktipembayarandata)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return buktipembayarandata, nil
	case nil:
		return buktipembayarandata, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return buktipembayarandata, err
}
