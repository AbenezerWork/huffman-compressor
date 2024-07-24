package main

import (
	"log"
	"os"
)

func decode() {
	node := decoded[0]
	file, err := os.ReadFile("bitdata.bin")
	log.Println("Trying to decode...")

	if err != nil {
		log.Fatal(err)
	}

	decodd := []byte{}
	curr := node

	log.Println(file[0:8])

	for i := range file {
		it := 0
		for it < 8 {
			if file[i]&(1<<it) != 0 {
				curr = decoded[curr.Right]
			} else {
				curr = decoded[curr.Left]
			}
			if curr.Chr != 0 {
				decodd = append(decodd, curr.Chr)
				curr = node
			}
			it += 1
		}
	}

	output, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	output.Write(decodd)
}
