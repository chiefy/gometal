package adapters

import (
	"fmt"
	"github.com/chiefy/gometal/gometal"
	_ "os"
	_ "text/template"
)

func ParseTemplates(gm *gometal.GoMetal) error {
	for idx, file := range gm.Files {
		fmt.Printf("%d\tname:%v\ttype:%v\n", idx, file.Info.Name(), file.Type.Extension)
	}
	return nil
}

func GoTemplate() gometal.Adapter {
	return ParseTemplates
}
