package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/popcount"
)

func main() {
	var (
		i   int
		err error
	)

	if len(os.Args) > 1 {
		i, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		i = 3
	}

	fmt.Println("i:", i)

	fmt.Printf("byte(%d)=%T(%d)=%03b\n", i, byte(i), i, i)
	fmt.Println("PopCount:", popcount.PopCount(uint64(i)))
	fmt.Println("PopCount2_3:", popcount.PopCount2_3(uint64(i)))
	fmt.Println("PopCount2_4:", popcount.PopCount2_4(uint64(i)))
	fmt.Println("PopCount2_5:", popcount.PopCount2_5(uint64(i)))
}

//!-
