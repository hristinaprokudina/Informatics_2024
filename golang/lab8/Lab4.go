package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func RunLab8A() {
	Xn, Xk, deltaX := readInputFile("input.txt")

	resultsA, xValuesA := taskA(Xn, Xk, deltaX)
	printResults("Значения Y для диапазона:", xValuesA, resultsA)

	extraValues := readExtraValues("input.txt")
	resultsB := taskB(extraValues)
	printResults("Значения Y для дополнительных значений:", extraValues, resultsB)
}

func readInputFile(filename string) (float64, float64, float64) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return 0, 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []float64

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err == nil {
			values = append(values, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}

	if len(values) < 3 {
		fmt.Println("Ошибка: недостаточно данных в файле.")
		return 0, 0, 0
	}

	return values[0], values[1], values[2]
}

func readExtraValues(filename string) []float64 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var extraValues []float64
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		if lineCount > 3 {
			value, err := strconv.ParseFloat(scanner.Text(), 64)
			if err == nil {
				extraValues = append(extraValues, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
	}

	return extraValues
}

func taskA(Xn float64, Xk float64, deltaX float64) ([]float64, []float64) {
	var results []float64
	var xValues []float64

	for x := Xn; x <= Xk; x += deltaX {
		y := calculateY(x)
		if !math.IsNaN(y) {
			results = append(results, y)
			xValues = append(xValues, x)
		}
	}
	return results, xValues
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

func printResults(header string, xValues []float64, results []float64) {
	fmt.Println(header)
	for i, y := range results {
		if i < len(xValues) {
			fmt.Printf("x = %.4f, Y = %.4f\n", xValues[i], y)
		} else {
			fmt.Printf("Y = %.4f\n", y)
		}
	}
}
