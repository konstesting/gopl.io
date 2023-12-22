// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
// !+
package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc[%d] = pc[%d/2] + byte(%d&1) = %d(%05b) + byte(%d) = %d(%05b)\n",
			i,
			i,
			i,
			pc[i/2],
			pc[i/2],
			i&1,
			pc[i],
			pc[i],
		)
	}
}

func PC(i int) byte {
	return pc[i]
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
Упражнение 2.3. Перепишите функцию PopCount так, чтобы она использовала
цикл вместо единого выражения. Сравните производительность двух версий. (В раз­
деле 11.4 показано, как правильно сравнивать производительность различных реали­
заций.)
*/
func PopCount2_3(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count = count + int(pc[byte(x>>(i*8))])
	}
	return count
}

/*
Упражнение 2.4. Напишите версию PopCount, которая подсчитывает биты с по­
мощью сдвига аргумента по всем 64 позициям, проверяя при каждом сдвиге крайний
справа бит. Сравните производительность этой версии с выборкой из таблицы.
*/
func PopCount2_4(x uint64) int {
	var count int
	for i := 1; i <= 64; i++ {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}

/*
Упражнение 2.5. Выражение х&(х-1) сбрасывает крайний справа ненулевой
бит х. Напишите версию PopCount, которая подсчитывает биты с использованием
этого факта, и оцените ее производительность
*/
func PopCount2_5(x uint64) int {
	var count int
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
