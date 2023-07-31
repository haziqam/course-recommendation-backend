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

func NewMatkul(namaMatkul string, SKS int, namaJurusan string, minSemester int, prediksiIndeks string) Matkul {
	return Matkul{
		NamaMatkul: namaMatkul,
		SKS: SKS,
		NamaJurusan: namaJurusan,
		MinSemester: minSemester,
		PrediksiIndeks: prediksiIndeks,
	}
}

func (matkul *Matkul) ScanRow(row *sql.Rows) error {
	return row.Scan(&matkul.NamaMatkul, &matkul.SKS, &matkul.NamaJurusan, &matkul.MinSemester, &matkul.PrediksiIndeks)
}

func (matkul Matkul) GetNilai() float32 {
	tabelNilai := map[string]float32{
		"A": 4.0,
		"AB": 3.5,
		"B": 3,
		"BC": 2.5,
		"C": 2,
		"D": 1,
		"E": 0,
	}
	return tabelNilai[matkul.PrediksiIndeks]
}

func CountTotalSKS(matkulArr []Matkul) int {
	totalSKS := 0
	for _, matkul := range matkulArr {
		totalSKS += matkul.SKS
	}
	return totalSKS
}

