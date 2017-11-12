package main

import (
	"os"
	"log"
	"fmt"
	"strings"
)

func main() {

	file, err := os.Open("./tmp/file")
	if (err != nil) {
		panic(err)
	}
	data := make([]byte, 20)
	count, err := file.Read(data)
	if (err != nil) {
		log.Fatal(err)
	}
	array :=strings.Split(string(data[:count]), " ")
	for index, str := range array{
		fmt.Printf("%dth word is %s\n", index, str)
	}



}

