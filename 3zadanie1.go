package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Введите строку:")

	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	simbols := make(map[string]int)
	var value int = 1

	sim := strings.Split(line, "")

	sort.Strings(sim)
	//fmt.Println(sim)

	var v1 string = ""
	for _, v := range sim {
		if v == " " {
			v = "\" \""
		}
		if v == v1 {
			simbols[v] = simbols[v] + 1
		} else {
			simbols[v] = value
		}
		v1 = v
	}

	delete(simbols, "\n")

	for k1, v1 := range simbols {
		fmt.Printf("Введённый символ: %v, количество повторений: %v\n", k1, v1)

	}
}
