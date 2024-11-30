package Lab4

import (
	"fmt"
	"math"
)

func RunLab4() {
	Xn := 0.11
	Xk := 0.36
	deltaX := 0.05
	taskA(Xn, Xk, deltaX)
	extraValues := []float64{0.2, 0.3, 0.38, 0.43, 0.57}
	taskB(extraValues)
}

func taskA(Xn, Xk, deltaX float64) [] float64 {
	fmt.Println("Значения Y для диапазона:")
	for x := Xn; x <= Xk; x += deltaX {
		y := calculateY(x)
		fmt.Printf("x = %.2f, Y = %.4f\n", x, y)
	}
}
func taskB(values []float64) []float64 {
	fmt.Println("\nЗначения Y для дополнительных значений:")
	for _, x := range values {
		y := calculateY(x)
		fmt.Printf("x = %.2f, Y = %.4f\n", x, y)
	}
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
