package gometal

import (
	"fmt"
	"github.com/chiefy/gometal/gometal/file"
	"os"
	"path/filepath"
)

type Adapter func(*GoMetal) error

type GoMetal struct {
	Files       []*file.GoMetalFile
	Middlewares []Adapter
	Metadata    map[string]string
	Source      string
	Dest        string
	Clean       bool
}

func MassageDirLocation(dirName string) (string, error) {
	var d string
	var err error

	if len(dirName) == 0 {
		return d, fmt.Errorf("no source directory given")
	}
	if filepath.IsAbs(dirName) {
		d = filepath.Clean(dirName)
	} else {
		d, err = filepath.Abs(dirName)
		if err != nil {
			return d, fmt.Errorf("bad relative path: '%v'", dirName)
		}
	}
	return d, nil
}

func NewGoMetal(sourceDir string, destDir string) (*GoMetal, error) {
	sd, err := MassageDirLocation(sourceDir)
	if err != nil {
		return nil, err
	}

	if len(destDir) == 0 {
		destDir = "dest"
	}
	dd, err := MassageDirLocation(destDir)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(sd); os.IsNotExist(err) {
		return nil, fmt.Errorf("source dir '%s' is inaccessible", sd)
	}
	// TODO: if destination dir doesn't exist, create it
	if _, err := os.Stat(dd); os.IsNotExist(err) {
		return nil, fmt.Errorf("destination dir '%s' is inaccessible", dd)
	}

	gm := &GoMetal{
		Source:      sd,
		Dest:        dd,
		Clean:       true,
		Middlewares: make([]Adapter, 0),
	}

	if gm.Files, err = file.LoadAllFiles(sd); err != nil {
		return nil, err
	}
	return gm, nil
}

func (gm *GoMetal) Use(mw Adapter) {
	gm.Middlewares = append(gm.Middlewares, mw)
}

func (gm *GoMetal) Build() error {
	for _, mf := range gm.Middlewares {
		if err := mf(gm); err != nil {

		}
	}
	return nil
}
