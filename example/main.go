package main

import (
	"fmt"
	"github.com/chiefy/gometal/adapters"
	"github.com/chiefy/gometal/gometal"
	"os"
)

func main() {
	gm, err := gometal.NewGoMetal("./src", "dest")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		os.Exit(1)
	}
	gm.Use(adapters.GoTemplate())
	gm.Build()
}
