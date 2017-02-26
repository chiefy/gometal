package gometal

import (
	"fmt"
	"os"
	"path/filepath"
)

type Adapter func(*GoMetal) error

type GoMetal struct {
	Files       []os.FileInfo
	Middlewares []Adapter
	Metadata    map[string]string
	Source      string
	Dest        string
	Clean       bool
}

func MassageSourceDir(sourceDir string) (string, error) {
	var sd string
	var err error

	if len(sourceDir) == 0 {
		return sd, fmt.Errorf("no source directory given")
	}
	if filepath.IsAbs(sourceDir) {
		sd = filepath.Clean(sourceDir)
	} else {
		sd, err = filepath.Abs(sourceDir)
		if err != nil {
			return sd, fmt.Errorf("bad relative path: '%v'", sourceDir)
		}
	}
	return sd, nil
}

func NewGoMetal(sourceDir string) (*GoMetal, error) {
	fixedDir, err := MassageSourceDir(sourceDir)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(fixedDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("Could not access directory '%s'", fixedDir)
	}
	gm := &GoMetal{
		Source:      fixedDir,
		Clean:       true,
		Middlewares: make([]Adapter, 0),
	}
	if err := gm.LoadFiles(); err != nil {
		return nil, fmt.Errorf("Could not read files in directory: '%v'", gm.Source)
	}
	return gm, nil
}

func (gm *GoMetal) LoadFiles() error {
	d, err := os.Open(gm.Source)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			gm.Files = append(gm.Files, fi)
		}
	}
	return nil
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
