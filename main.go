package main

import (
	"fmt"
	"time"
)

func main() {
	locations := []string{
		"Pacific/Enderbury",
		"Pacific/Kanton",
	}

	for _, l := range locations {
		location, err := time.LoadLocation(l)
		if err != nil {
			fmt.Println("💀 err: ", err)
			return
		}
		fmt.Println("✅ got: ", location)
	}

	fmt.Println("done")
}
