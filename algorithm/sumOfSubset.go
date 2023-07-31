package main

// import "fmt"

func findSubsetsWithSum(arr []int, minSum int, maxSum int) [][]int {
	// TODO: Add penjelasan
	arrayLength := len(arr)
	// Initialize dp table with all false values
	dp := make([][]bool, arrayLength + 1)
	for i := range dp {
		dp[i] = make([]bool, maxSum + 1)
	}

	// Base case: it is possible to obtain a sum of 0 using an empty subset
	for i := 0; i <= arrayLength; i++ {
		dp[i][0] = true
	}

	// Dynamic programming to fill the dp table
	for i := 1; i <= arrayLength; i++ {
		for j := 1; j <= maxSum; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i - 1][j] || dp[i - 1][j - arr[i - 1]]
			} else {
				dp[i][j] = dp[i - 1][j]
			}
		}
	}

	// Backtrack to find all subsets whose sum is maxSum
	var subsets [][]int

	for k := minSum; k <= maxSum; k++ {
		var currentSubset []int
		findSubsets(arr, arrayLength, k, dp, &currentSubset, &subsets)
	} 

	return subsets
}

func findSubsets(arr []int, i, sum int, dp [][]bool, currentSubset *[]int, subsets *[][]int) {
	//TODO: bikin penjelasan
	if i == 0 && sum != 0 && dp[0][sum] {
		// Add the current subset to the result
		*subsets = append(*subsets, append([]int{}, *currentSubset...))
		return
	}

	if i == 0 && sum == 0 {
		// Found a subset whose sum is K, add it to the result
		*subsets = append(*subsets, append([]int{}, *currentSubset...))
		return
	}

	if dp[i-1][sum] {
		// The current element is not included in the subset
		findSubsets(arr, i-1, sum, dp, currentSubset, subsets)
	}

	if sum >= arr[i-1] && dp[i-1][sum-arr[i-1]] {
		// The current element is included in the subset
		*currentSubset = append(*currentSubset, i - 1)
		findSubsets(arr, i-1, sum-arr[i-1], dp, currentSubset, subsets)
		*currentSubset = (*currentSubset)[:len(*currentSubset)-1] // backtrack
	}
}

// func main() {
// 	// arr := []int{1, 2, 3}
//  	arr := []int{100, 50, 45, 20, 10, 5, 20}
// 	// K := 50
// 	subsets := findSubsetsWithSum(arr, 40, 50)

// 	// Print all subsets whose sum is K
// 	for _, subset := range subsets {
// 		fmt.Println(subset)
// 	}
// }
