/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0. 1 1 100 100 1 100
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var x1, y1, x2, y2, x3, y3 float64
	fmt.Println("Введите координаты треугольника - x1, y1, x2, y2, x3, y3:")
	//fmt.Scanf("%d %d %d %d %d %d\n", &x1, &y1, &x2, &y2, &x3, &y3)
	fmt.Scanln(&x1)
	fmt.Scanln(&y1)
	fmt.Scanln(&x2)
	fmt.Scanln(&y2)
	fmt.Scanln(&x3)
	fmt.Scanln(&y3)

	b1 := proverka(x1, y1, x2, y2, x3, y3)
	fmt.Printf("Можно ли построить треугольник по заданным точкам: %t\n", b1)

	if b1 == true {
		gaus1 := ploshad(x1, y1, x2, y2, x3, y3)
		fmt.Printf("Площадь треугольника: %.2f ", gaus1)

		ugol1 := ugol(x1, y1, x2, y2, x3, y3)
		fmt.Printf("Является ли треугольник прямоугольным: %t\n", ugol1)
	}

}

func proverka(x1, y1, x2, y2, x3, y3 float64) bool {
	var b bool
	var gaus float64

	gaus = (x1-x3)*(y2-y3) - (x2-x3)*(y1-y3)
	if gaus == 0 {
		b = false
	} else {
		b = true
	}

	return b
}

func ploshad(x1, y1, x2, y2, x3, y3 float64) float64 {
	var gaus float64

	gaus = math.Abs(1.0 / 2.0 * ((x1-x3)*(y2-y3) - (x2-x3)*(y1-y3)))

	return gaus
}

func ugol(x1, y1, x2, y2, x3, y3 float64) bool {
	var ug bool
	var st1, st2, st3 float64

	st1 = math.Pow(math.Pow(x2-x1, 2)+math.Pow(y2-y1, 2), 1.0/2.0)
	st2 = math.Pow(math.Pow(x3-x1, 2)+math.Pow(y3-y1, 2), 1.0/2.0)
	st3 = math.Pow(math.Pow(x2-x3, 2)+math.Pow(y2-y3, 2), 1.0/2.0)

	for i := 1; i <= 3; i++ {
		if st1 > st2 {
			st1, st2 = st2, st1
		}
		if st2 > st3 {
			st2, st3 = st3, st2
		}
		if st1 > st3 {
			st1, st3 = st3, st1
		}
	}

	if (math.Pow(st1, 2) + math.Pow(st2, 2)) == math.Pow(st3, 2) {
		ug = true
	} else {
		ug = false
	}

	return ug
}
