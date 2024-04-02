package main

import (
	"fmt"
	"os"
)

/*
Упражнение 3.12. Напишите функцию, которая сообщает, являются ли две строки
анаграммами одна другой, т.е. состоят ли они из одних и тех же букв в другом порядке.
*/
func main() {
	if isAnagramma(os.Args[1], os.Args[2]) {
		fmt.Printf("The string '%s' is anagramma of '%s'\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("The string '%s' is NOT anagramma of '%s'\n", os.Args[1], os.Args[2])
	}
}

func isAnagramma(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s2len := len(s2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[s2len-1-i] {
			return false
		}
	}
	return true
}
