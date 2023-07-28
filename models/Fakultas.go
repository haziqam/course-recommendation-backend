package models

import (
	"database/sql"
)

type Fakultas struct {
	NamaFakultas string `json:"namaFakultas"`
}

func (fakultas *Fakultas) ScanRow(row *sql.Rows) error {
	return row.Scan(&fakultas.NamaFakultas)
}