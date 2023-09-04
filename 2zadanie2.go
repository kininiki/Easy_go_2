package main

import "fmt"

func main() {
	var d, l, pricezabor, zem float64
	fmt.Println("Введите ширину участка (м):")
	fmt.Scanf("%f\n", &d)
	fmt.Println("Введите длину участка (м):")
	fmt.Scanf("%f\n", &l)
	fmt.Println("Введите стоимость погонного метра забора (руб):")
	fmt.Scanf("%f\n", &pricezabor)
	fmt.Println("Введите стоимость кв. м. земли (руб):")
	fmt.Scanf("%f\n", &zem)

	p, s := dacha(l, d)
	pricez := zabor(p, pricezabor)
	z := uchastok(s, zem)

	fmt.Printf("Стоимость всей земли: %.2f руб.\n", z)
	fmt.Printf("Стоимость забора по периметру: %.2f руб.\n", pricez)

}

func dacha(l, d float64) (float64, float64) {
	p := l*2 + d*2
	s := l * d
	return p, s
}

func zabor(p, pricezabor float64) float64 {
	pricez := pricezabor * p
	return pricez
}

func uchastok(s, zem float64) float64 {
	z := s * zem * 1.3
	return z
}
