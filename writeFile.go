package main

import (
	"fmt"
	"log"
	"os"
)

func writeFile(table *map[byte][]bool, file *[]byte) {
	encoded := []bool{}

	for i := range *file {
		encoded = append(encoded, (*table)[(*file)[i]]...)
	}
	fmt.Println(encoded, len(encoded))
	curr := []bool{}
	encodedBytes := []byte{}
	for i := range encoded {
		curr = append(curr, encoded[i])
		if len(curr) == 8 {
			var num uint8
			for it := 0; it < 8; it++ {
				if curr[it] {
					num += (1 << it)
				}
			}
			encodedBytes = append(encodedBytes, byte(num))
			curr = []bool{}
		}
	}
	if len(curr) != 0 {
		num := byte(0)
		for it := 0; it < len(curr); it++ {
			if curr[it] {
				num += (1 << it)
			}
		}
		encodedBytes = append(encodedBytes, num)
	}
	fmt.Println("checkk", encodedBytes)
	out, err := os.Create("bitdata.bin")
	if err != nil {
		log.Fatal(err)
	}
	out.Write(encodedBytes)
}
