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
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DbInstance, err = sql.Open("postgres", connectionString)
	if err != nil {
			log.Fatal("Failed to connect to database. Error:\n", err)
			os.Exit(2)
	}
}


// rows, err := db.Query(`
// CREATE TABLE fakultas (
// 	nama_fakultas VARCHAR(100) PRIMARY KEY
// )
// `)

// fmt.Println(rows)
// CREATE TABLE jurusan (
// 	nama_jurusan	varchar(100) PRIMARY KEY,
// 	nama_fakultas 	varchar(100),
// 	CONSTRAINT fk_nama_fakultas FOREIGN KEY(nama_fakultas) REFERENCES fakultas(nama_fakultas)
// )

//CREATE TABLE matkul (
// 	nama_matkul		VARCHAR(100) PRIMARY KEY,
// 	sks				INT NOT NULL CHECK(sks > 0),
// 	nama_jurusan 	VARCHAR(100) REFERENCES jurusan(nama_jurusan),
// 	min_semester	INT NOT NULL CHECK(min_semester > 0),
// 	prediksi		VARCHAR(2) NOT NULL
// )

	// Check if the table "Fakultas" exists
	// var tableName string
	// err = DbInstance.QueryRow(`
	// 	SELECT table_name 
	// 	FROM information_schema.tables 
	// 	WHERE table_name = 'jurusan'
	// `).Scan(&tableName)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		fmt.Println("Table 'Jurusan' does not exist.")
	// 	} else {
	// 		log.Fatal("Failed to query database. Error:\n", err)
	// 		os.Exit(2)
	// 	}
	// } else {
	// 	fmt.Println("Table 'Jurusan' exists")
	// }