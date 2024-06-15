package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("go.mod")
	if err != nil {
		fmt.Println("Error creating go.mod:", err)
		return
	}
	defer f.Close()

	err = os.Mkdir("src", 0755)
	if err != nil {
		fmt.Println("Error creating src directory:", err)
		return
	}

	f, err = os.Create("src/main.go")
	if err != nil {
		fmt.Println("Error creating src/main.go:", err)
		return
	}
	defer f.Close()

	err = os.Mkdir("static", 0755)
	if err != nil {
		fmt.Println("Error creating static directory:", err)
		return
	}
}
