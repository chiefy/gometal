package file

import (
	"fmt"
	"gopkg.in/h2non/filetype.v1"
	"gopkg.in/h2non/filetype.v1/types"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
)

type GoMetalFile struct {
	Info os.FileInfo
	Type types.Type
	Data []byte
}

func GetFileType(data []byte, filePath string) (*types.Type, error) {
	fileType, err := filetype.Match(data)
	if err != nil {
		return nil, fmt.Errorf("couldn't match file type for %v - %v", filePath, err)
	}
	if fileType.Extension != string(types.Unknown.Extension) {
		return &fileType, nil
	}
	for {
		ext := filepath.Ext(filePath)
		fileType = filetype.GetType(ext)
		if fileType.Extension == string(types.Unknown.Extension) {
			mimetype := mime.TypeByExtension(ext)
			filetype.AddType(ext, mimetype)
		} else {
			break
		}
	}
	return &fileType, nil
}

func LoadFile(basePath string, fileInfo os.FileInfo) (*GoMetalFile, error) {
	filePath := basePath + string(os.PathSeparator) + fileInfo.Name()
	if !fileInfo.Mode().IsRegular() {
		return nil, fmt.Errorf("%v is not a regular file", filePath)
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read %v - %v", fileInfo.Name(), err)
	}
	fileType, err := GetFileType(data, filePath)
	if err != nil {
		return nil, err
	}
	gmf := &GoMetalFile{
		Info: fileInfo,
		Type: *fileType,
		Data: data,
	}
	return gmf, nil
}

func LoadAllFiles(sourceDir string) ([]*GoMetalFile, error) {
	d, err := os.Open(sourceDir)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}
	loadedFiles := make([]*GoMetalFile, 0)
	for _, fi := range fi {
		gmf, err := LoadFile(d.Name(), fi)
		if err != nil {
			return nil, err
		}
		loadedFiles = append(loadedFiles, gmf)
	}
	return loadedFiles, nil
}
