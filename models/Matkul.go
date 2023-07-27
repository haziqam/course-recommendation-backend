package models

type Matkul struct {
	NamaMatkul     string `json:"NamaMatkul"`
	SKS            int    `json:"sks"`
	NamaJurusan    string `json:"namaJurusan"`
	MinSemester    int    `json:"minSemester"`
	PrediksiIndeks string `json:"prediksiIndeks"`
}