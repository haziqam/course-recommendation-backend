package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DbInstance *sql.DB

func ConnectDB() {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DbInstance, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("Failed to connect to database. Error:\n", err)
	}

	initializeTables(DbInstance)
}

func initializeTables(DbInstance *sql.DB) {
	_, err := DbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS fakultas (
			nama_fakultas character varying(100) NOT NULL,
			CONSTRAINT fakultas_pkey PRIMARY KEY (nama_fakultas)
		)
	`)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = DbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS jurusan (
			nama_jurusan character varying(100) NOT NULL,
			nama_fakultas character varying(100),
			CONSTRAINT jurusan_pkey PRIMARY KEY (nama_jurusan),
			CONSTRAINT fk_nama_fakultas FOREIGN KEY (nama_fakultas)
				REFERENCES public.fakultas (nama_fakultas) MATCH SIMPLE
				ON UPDATE CASCADE
				ON DELETE CASCADE
		)
	`)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = DbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS matkul (
			nama_matkul character varying(100) NOT NULL,
			sks integer NOT NULL,
			nama_jurusan character varying(100) NOT NULL,
			min_semester integer NOT NULL,
			prediksi character varying(2) NOT NULL,
			CONSTRAINT matkul_pkey PRIMARY KEY (nama_matkul),
			CONSTRAINT matkul_sks_check CHECK (sks > 0),
			CONSTRAINT matkul_min_semester_check CHECK (min_semester > 0),
			CONSTRAINT fk_nama_jurusan FOREIGN KEY (nama_jurusan)
				REFERENCES public.jurusan (nama_jurusan) MATCH SIMPLE
				ON UPDATE CASCADE
				ON DELETE CASCADE
		)
	`)

	if err != nil {
		log.Fatalln(err)
	}
}
