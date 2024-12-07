package Lab4

import (
	"fmt"
	"math"
)

func RunLab4() {
	Xn := 0.11
	Xk := 0.36
	deltaX := 0.05
	
	resultsA := taskA(Xn, Xk, deltaX)
	printResults("Значения Y для диапазона:", resultsA)

	extraValues := []float64{0.2, 0.3, 0.38, 0.43, 0.57}
	resultsB := taskB(extraValues)
	printResults("Значения Y для дополнительных значений:", resultsB)
}
func taskA(Xn float64, Xk float64, deltaX float64) {
	fmt.Println("Значения Y для диапазона:")
	for x := Xn; x <= Xk; x += deltaX {
		y := calculateY(x)
		if !math.IsNaN(y) {
			fmt.Printf("x = %.2f, Y = %.4f\n", x, y)
		}
	}
	return results
}

func taskB(values []float64) []float64 {
	var results []float64
	for _, x := range values {
		y := calculateY(x)
		if !math.IsNaN(y) {
			results = append(results, y)
		}
	}
	return results
}

func calculateY(x float64) float64 {
	if x <= 0 {
		fmt.Printf("Ошибка: логарифм не определен для x = %f\n", x)
		return math.NaN()
	}
	sinCubed := math.Pow(math.Sin(x), 3)
	cosCubed := math.Pow(math.Cos(x), 3)
	return (sinCubed + cosCubed) * math.Log(x)
}
func printResults(header string, results []float64) {
	fmt.Println(header)
	for i, y := range results {
		fmt.Printf("x = %.2f, Y = %.4f\n", 0.2+float64(i)*0.1, y)
	}
}
