/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Введите размер шахматной доски (от 0 до 20): ")
	fmt.Scanf("%d\n", &n)

	if n >= 0 && n <= 20 {
		s := make([][]int, n)
		//		k1, k2 := 0, 1
		var a int = 0

		for i := range s {
			s[i] = make([]int, n)
		}

		for j := 0; j < n; j = j + 1 {
			for i := 0; i < n; i = i + 1 {
				s[i][j] = 1
			}
		}

		for j := 0; j < n; j = j + 1 {
			if a == 0 {
				for i := 0; i < n; i = i + 2 {
					s[i][j] = 0
				}
				a = 1
			} else if a == 1 {
				for i := 1; i < n; i = i + 2 {
					s[i][j] = 0
				}
				a = 0
			}

		}

		for i := range s {
			fmt.Println(s[i])
		}

		if n == 0 {
			fmt.Printf("0 \n")
		}

	} else {
		fmt.Printf("Вы ввели некорректное значение.")
	}

}
