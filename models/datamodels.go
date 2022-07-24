package models

import (
	"apiyeah/config"
	"database/sql"
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
	Id_tipe_pembayarans       int64  `json:"id_tipe_pembayarans"`
	Nama_tipe_pembayarans     string `json:"nama_tipe_pembayarans"`
	Id_pembayarans            int64  `json:"id_pembayarans"`
	Nama_pembayarans          string `json:"nama_pembayarans"`
	Nama_rekening_pembayarans string `json:"nama_rekening_pembayarans"`
	No_rekening_pembayarans   string `json:"no_rekening_pembayarans"`
	Logo_Pembayarans          string `json:"logo_pembayarans"`
}

type TipePembayaran struct {
	Id_tipe_pembayarans   int64  `json:"id_tipe_pembayarans"`
	Nama_tipe_pembayarans string `json:"nama_tipe_pembayarans"`
}

type StatusPembayaran struct {
	Id_status_pembayarans   int64  `json:"id_status_pembayarans"`
	Nama_status_pembayarans string `json:"nama_status_pembayarans"`
}

type SosialMedia struct {
	Id_sosial_media   int64  `json:"id_sosial_medias"`
	Nama_sosial_media string `json:"nama_sosial_medias"`
	Icon_sosial_media string `json:"icon_sosial_medias"`
	Url_sosial_media  string `json:"url_sosial_medias"`
}

type MetaTag struct {
	Id_meta_tag     int64  `json:"id_meta_tags"`
	Nama_meta_tag   string `json:"nama_meta_tags"`
	Konten_meta_tag string `json:"konten_meta_tags"`
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
						id_tipe_pembayarans,
						nama_tipe_pembayarans,
						id_pembayarans,
						nama_pembayarans,
						nama_rekening_pembayarans,
						no_rekening_pembayarans,
						logo_pembayarans
					FROM master_pembayarans
					JOIN master_tipe_pembayarans ON master_tipe_pembayarans.id_tipe_pembayarans=master_pembayarans.tipe_pembayarans_id
					WHERE status_hapus_pembayarans=false`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var pembayaran Pembayaran

		err = rows.Scan(&pembayaran.Id_tipe_pembayarans,
			&pembayaran.Nama_tipe_pembayarans,
			&pembayaran.Id_pembayarans,
			&pembayaran.Nama_pembayarans,
			&pembayaran.Nama_rekening_pembayarans,
			&pembayaran.No_rekening_pembayarans,
			&pembayaran.Logo_Pembayarans,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		pembayarans = append(pembayarans, pembayaran)
	}

	return pembayarans, err
}

func AmbilSemuaTipePembayaran() ([]TipePembayaran, error) {
	db := config.CreateConnection()

	defer db.Close()

	var tipepembayarans []TipePembayaran

	sqlStatement := `SELECT
						id_tipe_pembayarans,
						nama_tipe_pembayarans
					FROM master_tipe_pembayarans`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var tipepembayaran TipePembayaran

		err = rows.Scan(&tipepembayaran.Id_tipe_pembayarans,
			&tipepembayaran.Nama_tipe_pembayarans,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		tipepembayarans = append(tipepembayarans, tipepembayaran)
	}

	return tipepembayarans, err
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

func AmbilSemuaSosialMedia() ([]SosialMedia, error) {
	db := config.CreateConnection()

	defer db.Close()

	var sosialmedias []SosialMedia

	sqlStatement := `SELECT
						id_sosial_medias,
						nama_sosial_medias,
						icon_sosial_medias,
						url_sosial_medias
					FROM master_sosial_medias`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var sosialmedia SosialMedia

		err = rows.Scan(&sosialmedia.Id_sosial_media,
			&sosialmedia.Nama_sosial_media,
			&sosialmedia.Icon_sosial_media,
			&sosialmedia.Url_sosial_media,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		sosialmedias = append(sosialmedias, sosialmedia)
	}

	return sosialmedias, err
}

func AmbilSemuaMetaTag() ([]MetaTag, error) {
	db := config.CreateConnection()

	defer db.Close()

	var metatags []MetaTag

	sqlStatement := `SELECT
						id_meta_tags,
						nama_meta_tags,
						regexp_replace(coalesce(konten_meta_tags, ''), E'[\\n\\r]+', ' ', 'g' ) AS konten_meta_tags
					FROM master_meta_tags`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var metatag MetaTag

		err = rows.Scan(&metatag.Id_meta_tag,
			&metatag.Nama_meta_tag,
			&metatag.Konten_meta_tag,
		)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		metatags = append(metatags, metatag)
	}

	return metatags, err
}
