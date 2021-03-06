package serial

import (
	"fmt"
)

func PrintHex(value interface{}) {
	fmt.Printf("0x%02X\n", value)
}

func PrintBin(value interface{}) {
	fmt.Printf("0b%b\n", value)
}

func PrintData(data Data) {
	println("       |  0  1  2  3   4  5  6  7   8  9  A  B   C  D  E  F")
	fmt.Printf("%06x | ", 0)
	i := 1
	for ; i <= len(data); i++ {
		fmt.Printf("%02X ", data[i - 1])

		if i % 4 == 0 {
			print(" ")
		}
		if i % 16 == 0 {
			fmt.Printf("\n%06x | ", i)
		}
	}

	for {
		if (i - 1) % 16 == 0 {
			break
		}
		print("-- ")
		if i % 4 == 0 {
			print(" ")
		}
		i++
	}

	println()
}