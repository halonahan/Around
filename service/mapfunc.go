package main

import "fmt"

func main() {

	m := map[string]bool {
		"breakfast": true,
		"lunch": false,
		"dinner": false,
	}
	for key, value := range m {
		fmt.Println("key: ", key, " value: ", value)
	}

}
