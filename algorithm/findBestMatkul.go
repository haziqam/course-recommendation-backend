package algorithm

import (
	// "fmt"

	"github.com/haziqam/course-scheduler-backend/packages/models"
)

// returns []Matkul, IP, SKS
// assumes totalSKS di availables >= minSKS and maxSKS >= minSKS
func FindBestMatkul(availables []models.Matkul, minSKS, maxSKS int) ([]models.Matkul, float32, int) {
	var SKS []int
	var prediksiNilai []float32

	for _, matkul := range availables {
		SKS = append(SKS, matkul.SKS)
		prediksiNilai = append(prediksiNilai, matkul.GetNilai())
	}

	feasibleChoices := findSubsetsWithSum(SKS, minSKS, maxSKS)
	var highestIP float32 = 0.0
	var bestChoiceIndex int
	var bestChoiceSKS int = 0

	for i, choice := range feasibleChoices {
		currentIP, currentSKS := countIPandSKS(SKS, prediksiNilai, choice)
		if currentIP > highestIP {
			highestIP = currentIP
			bestChoiceIndex = i
			bestChoiceSKS = currentSKS
		} else if currentIP == highestIP && currentSKS > bestChoiceSKS {
			bestChoiceIndex = i
			bestChoiceSKS = currentSKS
		}
	}

	bestChoice := feasibleChoices[bestChoiceIndex]

	var bestMatkuls []models.Matkul
	for _, index := range bestChoice {
		bestMatkuls = append(bestMatkuls, availables[index])
	}

	return bestMatkuls, highestIP, bestChoiceSKS

}

func countIPandSKS(SKS []int, prediksiNilai []float32, choice []int) (float32, int) {
	var sumNilai float32 = 0.0
	var sumSKS int = 0
	for _, index := range choice {
		sumNilai += (float32(SKS[index]) * prediksiNilai[index])
		sumSKS += SKS[index]
	}

	return sumNilai / float32(sumSKS), sumSKS
}

// func main() {
// 	// availables := []models.Matkul{
// 	// 	models.NewMatkul("m1", 2, "IF", 1, "A"),
// 	// 	models.NewMatkul("m2", 1, "IF", 1, "A"),
// 	// 	models.NewMatkul("m3", 3, "IF", 1, "B"),
// 	// 	models.NewMatkul("m4", 1, "IF", 1, "BC"),
// 	// 	models.NewMatkul("m5", 4, "IF", 1, "C"),
// 	// 	models.NewMatkul("m6", 1, "IF", 1, "D"),
// 	// 	models.NewMatkul("m7", 2, "IF", 1, "E"),
// 	// 	models.NewMatkul("m8", 3, "IF", 1, "AB"),
// 	// 	models.NewMatkul("m9", 4, "IF", 1, "AB"),
// 	// 	models.NewMatkul("m10", 1, "IF", 1, "B"),
// 	// }
// 	availables := []models.Matkul{
// 		models.NewMatkul("m1", 2, "IF", 1, "A"),
// 		models.NewMatkul("m2", 1, "IF", 1, "A"),
// 		models.NewMatkul("m3", 3, "IF", 1, "A"),
// 		models.NewMatkul("m4", 4, "IF", 1, "A"),
// 		models.NewMatkul("m5", 1, "IF", 1, "A"),
// 		models.NewMatkul("m6", 1, "IF", 1, "A"),
// 		models.NewMatkul("m7", 2, "IF", 1, "A"),
// 		models.NewMatkul("m8", 3, "IF", 1, "A"),
// 		models.NewMatkul("m9", 4, "IF", 1, "A"),
// 		models.NewMatkul("m10", 1, "IF", 1, "A"),
// 	}

// 	fmt.Println(findBestMatkul(availables, 22, 22))
// }
