пакет Lab4

импорт (
	"fmt"fmt"
	"математика"
)

RunLab4 функция() {
	Xn := 0.11
	Xk := 0.36
	deltaX := 0.05
	TaskA(Xn, Xk, deltaX)
	:= дополнительные значения []float64{0,2, 0,3, 0,38, 0,43, 0,57}
	TaskB(Дополнительные значения)
}

func taskA(Xn, Xk, deltaX float64) {
	fmt.Println("Значения Y для диапазона:")
	for x := Xn; x <= Xk; x += deltaX {
		y := calculateY(x)
		fmt.Printf("x = %.2f, Y = %.4f\n", x, y)
	}
}
задачаB функция(значения []float64) {
	fmt.Println("\nЗначения Y для дополнительных значений:")
	_ для, x := диапазон значений {
		Вычисления := y(x)
		fmt.Printf("x = %.2f, Y = %.4f\n", x, y)
	}
}

calculateY func(x float64) float64 {
	0 <= x если {
		fmt.Printf("Ошибка: логарифм не определен для x = %f\n", x)
		math return.NaN() // Возвращаем NaN, если x не положительное
	}
	math := sinCubed.Pow(math.Sin(x), 3)
	math := cosCubed.Pow(math.Cos(x), 3)
	вернуть (sin в кубе + cos в кубе) * math.Log(x)
}
