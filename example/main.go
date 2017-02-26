package main

import (
	"fmt"
	"github.com/chiefy/gometal/gometal"
	"os"
)

func drafts() gometal.Adapter {
	var mw gometal.Adapter
	mw = func(gm *gometal.GoMetal) error {
		fmt.Printf("Files: %v", gm.Files)
		return nil
	}
	return mw
}

func main() {
	gm, err := gometal.NewGoMetal("./src")

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		os.Exit(1)
	}

	gm.Use(drafts())
	gm.Build()

}
