package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func setHeader() {
	file, err := os.Create("header.dat")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(nodes)
	if err != nil {
		log.Fatal(err)
	}
}

var decoded map[uint64]Node

func getHeader(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&decoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("decoded", decoded)
}
