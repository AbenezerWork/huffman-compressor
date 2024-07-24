package main

import (
	"flag"
	"log"
	"os"
)

var Cnt map[byte]int64

func Count(name string) []byte {
	Cnt = make(map[byte]int64)

	file, err := os.ReadFile(name)

	if err != nil {
		log.Panic(err)
	}
	for i := range file {
		Cnt[file[i]]++
	}

	return file
}

var output *string
var input *string

func init() {
	output = flag.String("o", "", "")
	input = flag.String("i", "", "")
}

func main() {
	flag.Parse()
	if *input != "" {
		file := Count(*input)
		log.Println(len(file))
		writeFile(buildTree(), &file)
		setHeader()
	} else if *output != "" {
		getHeader(*output)
		decode()
	}
}
