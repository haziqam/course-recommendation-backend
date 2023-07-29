package main

import "fmt"

func main() {
	// TC 1
	// const jumlahMatkul = 4
	// const maxSKSdiambil = 16

	// var SKS [jumlahMatkul]int = [jumlahMatkul]int{2, 5, 10, 5}
	// var nilai [jumlahMatkul]int = [jumlahMatkul]int{20, 30, 50, 10}
	// var dpTable [jumlahMatkul + 1][maxSKSdiambil + 1]int
	// var keputusan [jumlahMatkul + 1][maxSKSdiambil + 1][jumlahMatkul]int

	// TC 3
	// const jumlahMatkul = 4
	// const maxSKSdiambil = 16

	// var SKS [jumlahMatkul]int = [jumlahMatkul]int{6, 5, 10, 5}
	// var nilai [jumlahMatkul]int = [jumlahMatkul]int{12, 15, 50, 10}
	// var dpTable [jumlahMatkul + 1][maxSKSdiambil + 1]int
	// var keputusan [jumlahMatkul + 1][maxSKSdiambil + 1][jumlahMatkul]int

	// TC 4
	// const jumlahMatkul = 6
	// const maxSKSdiambil = 100

	// var SKS [jumlahMatkul]int = [jumlahMatkul]int{100, 50, 45, 20, 10, 5}
	// var nilai [jumlahMatkul]int = [jumlahMatkul]int{40, 35, 18, 4, 10, 2}
	// var dpTable [jumlahMatkul + 1][maxSKSdiambil + 1]int
	// var keputusan [jumlahMatkul + 1][maxSKSdiambil + 1][jumlahMatkul]int

	// TC 2
	const jumlahMatkul = 3
	const maxSKSdiambil = 5

	var SKS [jumlahMatkul]int = [jumlahMatkul]int{2, 1, 3}
	var nilai [jumlahMatkul]float32 = [jumlahMatkul]float32{65.0, 80.0, 30.0}
	var dpTable [jumlahMatkul + 1][maxSKSdiambil + 1]float32
	var keputusan [jumlahMatkul + 1][maxSKSdiambil + 1][jumlahMatkul]int
	for tahap := 1; tahap <= jumlahMatkul; tahap++ {
		for sisaSKS := 0; sisaSKS <= maxSKSdiambil; sisaSKS++ {
			currentMatkul := tahap - 1
			if sisaSKS-SKS[currentMatkul] < 0 {
				// Case 1: matkul pada tahap ini tidak diambil karena sisa SKS tidak mencukupi atau
				// karena mengambil matkul tidak lebih menguntungkan dibandingkan tidak mengambil
				copy(keputusan[tahap][sisaSKS][:], keputusan[tahap - 1][sisaSKS][:])
				dpTable[tahap][sisaSKS] = dpTable[tahap-1][sisaSKS]
			} else{
				nilaiJikaMengambil := countJlhNilai(nilai[:], SKS[:], keputusan[tahap - 1][sisaSKS-SKS[currentMatkul]][:]) + (nilai[currentMatkul] * float32(SKS[currentMatkul]))
				nilaiJikaMengambil /= (float32(countSKSdiambil(nilai[:], SKS[:], keputusan[tahap - 1][sisaSKS-SKS[currentMatkul]][:])) + float32(SKS[currentMatkul]))
				if nilaiJikaMengambil < dpTable[tahap - 1][sisaSKS] {
					// Case 2: matkul pada tahap ini tidak diambil karena mengambil
					// matkul tidak lebih menguntungkan dibandingkan tidak mengambil
					copy(keputusan[tahap][sisaSKS][:], keputusan[tahap - 1][sisaSKS][:])
					dpTable[tahap][sisaSKS] = dpTable[tahap-1][sisaSKS]
				} else {
					// Case 3: matkul pada tahap ini diambil karena sisa SKS mencukupi dan 
					// mengambil matkul lebih menguntungkan dibandingkan tidak mengambil
					copy(keputusan[tahap][sisaSKS][:], keputusan[tahap][sisaSKS-SKS[currentMatkul]][:])
					keputusan[tahap][sisaSKS][currentMatkul] = 1
					dpTable[tahap][sisaSKS] = countIP(nilai[:], SKS[:], keputusan[tahap][sisaSKS][:])
				}
			} 
			
		}
	}

	for _, row := range dpTable {
		fmt.Println(row)
	}

	for _, row := range keputusan {
		fmt.Println(row)
	}
}

func countIP(nilai []float32, SKS []int, keputusan []int) float32 {
	return countJlhNilai(nilai, SKS, keputusan) / float32(countSKSdiambil(nilai, SKS, keputusan))
}

func countJlhNilai(nilai []float32, SKS []int, keputusan []int) float32 {
	var jlhNilai float32 = 0.0
	banyakMatkul := len(nilai)

	for i := 0; i < banyakMatkul; i++ {
		jlhNilai += float32(keputusan[i]) * nilai[i] * float32(SKS[i])
	}
	// fmt.Println(jlhNilai, SKSdiambil)
	return jlhNilai
}

func countSKSdiambil (nilai []float32, SKS []int, keputusan []int) int {
	var SKSdiambil float32 = 0
	banyakMatkul := len(nilai)

	for i := 0; i < banyakMatkul; i++ {
		SKSdiambil += float32(keputusan[i]) * float32(SKS[i])
	}
	// fmt.Println(jlhNilai, SKSdiambil)
	return int(SKSdiambil)
}



// func main() {
// 	decision := []int{1, 0, 1}
// 	nilai := []float32{3.5, 2.0, 4.0}
// 	SKS := []int{3, 2, 4}
// 	fmt.Println(countIP(nilai, SKS, decision))
	
// }