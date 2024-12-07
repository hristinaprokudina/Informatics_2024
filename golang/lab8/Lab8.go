package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunLab8() {
	filename := "input.txt"

	createAndWriteFile(filename)

	readFile(filename)

	processFileData(filename)

	var searchTerm string
	fmt.Print("Введите текст для поиска в файле: ")
	fmt.Scanln(&searchTerm)
	searchInFile(filename, searchTerm)
}

func createAndWriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Println("Введите значения для записи в файл (введите 'exit' для завершения):")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		writer.WriteString(text + "\n")
	}

	writer.Flush()
	fmt.Println("Данные успешно записаны в файл:", filename)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Содержимое файла:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

func processFileData(filename string) {
	file, err := os.Open(filename)
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
	taskA(a, b, deltaX) 
	resultsB := taskB(xValues) 
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

// Функция для поиска текста в файле
func searchInFile(filename string, searchTerm string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchTerm) {
			fmt.Println("Найдено:", line)
			found = true
		}
	}

	if !found {
		fmt.Println("Текст не найден в файле.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
