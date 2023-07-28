package models

import (
	"database/sql"
)

type Jurusan struct {
	NamaJurusan  string `json:"namaJurusan"`
	NamaFakultas string `json:"namaFakultas"`
}

func (jurusan *Jurusan) ScanRow(row *sql.Rows) error {
	return row.Scan(&jurusan.NamaJurusan, &jurusan.NamaFakultas)
}