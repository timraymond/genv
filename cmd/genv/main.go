package main

import (
	"fmt"
	"os"

	"github.com/timraymond/genv/cmd/genv/internal"
)

func main() {
	if err := internal.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
