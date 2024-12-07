package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func RunLab8A() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var values []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Ошибка при чтении числа:", err)
			continue
		}
		values = append(values, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	if len(values) < 2 {
		fmt.Println("Недостаточно значений в файле")
		return
	}

	a := values[0]
	b := values[1]

	xValues := values[2:]

	deltaX := 0.05 
	taskAA(a, b, deltaX) 
	resultsB := taskBA(xValues) 
	printResultsA("Значения Y для дополнительных значений:", resultsB) 
}

func taskAA(Xn float64, Xk float64, deltaX float64) {
	fmt.Println("Значения Y для диапазона:")
	for x := Xn; x <= Xk; x += deltaX {
		y := calculateY(x)
		if !math.IsNaN(y) {
			fmt.Printf("x = %.2f, Y = %.4f\n", x, y)
		}
	}
}

func taskBA(values []float64) []float64 {
	var results []float64
	for _, x := range values {
		y := calculateY(x)
		if !math.IsNaN(y) {
			results = append(results, y)
		}
	}
	return results
}

func calculateYA(x float64) float64 {
	if x <= 0 {
		fmt.Printf("Ошибка: логарифм не определен для x = %f\n", x)
		return math.NaN()
	}
	sinCubed := math.Pow(math.Sin(x), 3)
	cosCubed := math.Pow(math.Cos(x), 3)
	return (sinCubed + cosCubed) * math.Log(x)
}

func printResultsA(header string, results []float64) {
	fmt.Println(header)
	for i, y := range results {
		fmt.Printf("x = %.2f, Y = %.4f\n", 0.2+float64(i)*0.1, y)
	}
}
