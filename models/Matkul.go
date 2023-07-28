package models

import (
	"database/sql"
)

type Matkul struct {
	NamaMatkul     string `json:"namaMatkul"`
	SKS            int    `json:"sks"`
	NamaJurusan    string `json:"namaJurusan"`
	MinSemester    int    `json:"minSemester"`
	PrediksiIndeks string `json:"prediksiIndeks"`
}

func (matkul *Matkul) ScanRow(row *sql.Rows) error {
	return row.Scan(&matkul.NamaMatkul, &matkul.SKS, &matkul.NamaJurusan, &matkul.MinSemester, &matkul.PrediksiIndeks)
}

