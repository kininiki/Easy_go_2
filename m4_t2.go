/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

0,5,6,7,8,9 - ок; 1 - ка; 2,3,4 - ки
10-19 - ок;

In: 22
Out: 22 бутылки
*/

package main

import "fmt"

func main() {

	var n int
	var end1 string
	fmt.Printf("Введите количество бутылок:")
	fmt.Scanf("%d", &n)

	if n >= 0 && n <= 200 {
		end1 = ending(n)
		fmt.Printf("%d бутыл%s ", n, end1)

	} else {
		fmt.Printf("Вы ввели некорректное значение.")
	}

}

func ending(n int) string {
	var end string
	var n1 int

	if n >= 20 && n <= 99 {
		n1 = n % 10
	} else if (n >= 0 && n <= 19) || (n >= 111 && n <= 119) {
		n1 = n
	} else {
		n1 = (n % 10) % 10
	}

	ok := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 111, 112, 113, 114, 115, 116, 117, 118, 119}
	ki := []int{2, 3, 4}
	ka := []int{1}

	for _, value := range ok {
		if value == n1 {
			end = "ок"
		}
	}
	for _, value := range ki {
		if value == n1 {
			end = "ки"
		}
	}
	for _, value := range ka {
		if value == n1 {
			end = "ка"
		}
	}

	return end
}
