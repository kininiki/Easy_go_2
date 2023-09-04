/*
Вход: слово(без пробела)
Выход: количество буквы 'a'(латинская) в слове. Если нет, то написать об этом.
*/

package main

import "fmt"

func main() {
	var str string
	var num int = 0

	fmt.Println("Введите слово на латинском (без пробела): ")
	fmt.Scanf("%s\n", &str)

	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			num++
		}
	}

	if num > 0 {
		fmt.Println("Количество букв \"a\" в слове равно:", num)
	} else {
		fmt.Println("В слове нет букв \"a\".")
	}

}
