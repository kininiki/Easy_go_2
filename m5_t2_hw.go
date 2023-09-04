/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/

package main

import (
	"fmt"
)

func main() {

	var code string
	fmt.Println("Введите строку из цифр от 00 до 26:")
	fmt.Scanf("%s\n", &code)

	var str1 string = codeToString(code)
	fmt.Printf("Вы ввели текст: %s\n", str1)

}

func codeToString(code string) string {
	var str string = ""
	st := make([]string, 100)

	mapa := map[string]string{"00": "a", "01": "b", "02": "c", "03": "d", "04": "e", "05": "f", "06": "g", "07": "h", "08": "i", "09": "j", "10": "k", "11": "l", "12": "m", "13": "n", "14": "o", "15": "p", "16": "q", "17": "r", "18": "s", "19": "t", "20": "u", "21": "v", "22": "w", "23": "x", "24": "y", "25": "z", "26": " "}
	s := code

	for i := 0; i < len(s)-1; i = i + 2 {
		k := string(rune(s[i]))
		k1 := string(rune((s[i+1])))
		st[i] = k
		st[i] += k1
		str2 := str

		for key, _ := range mapa {
			if key == st[i] {
				str = str + mapa[key]
			}
		}

		if str == str2 {
			str = "Вы ввели значение, которого нет в алфавите!"
			break
		}

		if len(s)%2 == 1 {
			str = "Вы ввели нечётное количество цифр!"
		}
	}

	return str
}
