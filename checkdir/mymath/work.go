package mymath

import (
	"net/http"
)

func CalcSumAndAvg(numbers []int) (int, float64) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum / len(numbers))
	return sum, average
}

func calculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateAverage(numbers []int) float64 {
	sum := calculateSum(numbers)
	average := float64(sum) / float64(len(numbers))
	return average
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
