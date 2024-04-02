// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		//fmt.Printf("  %s\n", comma(os.Args[i]))
		//fmt.Printf("  %s\n", comma3_10(os.Args[i]))
		fmt.Printf("  %s\n", comma3_11(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
/*
Упражнение 3.10. Напишите нерекурсивную версию функции comma, использующую
bytes.Buffer вместо конкатенации строк.
*/
func comma3_10(s string) string {
	var buf bytes.Buffer
	k := len(s) % 3
	if k == 0 {
		k = 3
	}
	for i := k; i <= len(s); i = i + 3 {
		if i-3 < 0 {
			fmt.Fprintf(&buf, "%s", s[:i])
		} else {
			fmt.Fprintf(&buf, "%s", s[i-3:i])
		}
		if i < len(s) {
			fmt.Fprintf(&buf, ",")
		}
	}
	return buf.String()
}

/*
Упражнение 3.11. Усовершенствуйте функцию comma так, чтобы она корректно
работала с числами с плавающей точкой и необязательным знаком.
*/
func comma3_11(s string) string {
	s = strings.ReplaceAll(s, ",", "")
	ss := strings.Split(s, ".")
	if len(ss) == 2 {
		return comma(ss[0]) + "." + comma(ss[1])
	}
	if len(ss) > 2 {
		log.Fatal("incorrect value: ", ss)
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
