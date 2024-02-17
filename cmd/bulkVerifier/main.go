package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bulkVerify <input.csv> <output.csv>")
		os.Exit(1)
	}

	BulkVerify(os.Args[1], os.Args[2])
}
