package main

import (
 "fmt"
 "math"
)

func calculateY(x float64) (float64, error) {
 if x <= 0 {
  return 0, fmt.Errorf("значение Y не определено (логарифм от неположительного числа)")
 }
 y := (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
 return y, nil
}

func main() {
 xn := 0.11
 xk := 0.36
 deltaX := 0.05

 fmt.Println("Задание A:")
 for x := xn; x <= xk; x += deltaX {
  y, err := calculateY(x)
  if err != nil {
   fmt.Printf("x = %.2f: %s\n", x, err)
  } else {
   fmt.Printf("x = %.2f: Y = %.6f\n", x, y)
  }
 }

 xValues := []float64{0.2, 0.3, 0.38, 0.43, 0.57}

 fmt.Println("\nЗадание B:")
 for _, x := range xValues {
  y, err := calculateY(x)
  if err != nil {
   fmt.Printf("x = %.2f: %s\n", x, err)
  } else {
   fmt.Printf("x = %.2f: Y = %.6f\n", x, y)
  }
 }
}