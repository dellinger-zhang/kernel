package fs

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File object.
type File struct {
	Folder      string
	FileName    string
	FileRawName string
	FileExtName string
	FullPath    string
}

// NewFile return parsed file object
func NewFile(path string) *File {

	file := new(File)
	file.FullPath = path
	file.Folder = filepath.Dir(path)
	file.FileName = filepath.Base(path)
	file.FileExtName = filepath.Ext(path)
	if file.FileExtName == "" {
		file.FileRawName = file.FileName
	} else {
		file.FileRawName = strings.Replace(file.FileName, file.FileExtName, "", -1)
	}

	return file
}

// Exists returns file exists or not
func (f *File) Exists() bool {
	return FileExists(f.FullPath)
}

// Delete file
func (f *File) Delete() bool {
	return RemoveFile(f.FullPath)
}

// Create a file.
func (f *File) Create() error {

	if !FileExists(f.Folder) {
		MakeDirP(f.Folder)
	}

	if _, err := os.Create(f.FullPath); err != nil {
		return err
	}

	return nil
}

// ReadAll text in file.
func (f *File) ReadAll() (string, error) {

	if !FileExists(f.FullPath) {
		return "", os.ErrNotExist
	}

	buf, err := ioutil.ReadFile(f.FullPath)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (f *File) Append(content []byte) error {
	file, err := os.OpenFile(f.FullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, ReadWriteMode)
	if err != nil {
		return err
	}

	n, err := file.Write(content)
	if err == nil && n < len(content) {
		err = io.ErrShortWrite
	}

	if err1 := file.Close(); err == nil {
		err = err1
	}
	return nil
}

// WriteAll text to file.
func (f *File) WriteAll(content string) error {

	return ioutil.WriteFile(f.FullPath, []byte(content), ReadWriteMode)
}

// ClearAll truncate whole file.
func (f *File) ClearAll() error {

	if !f.Delete() {
		return errors.New("fail to clear file")
	}
	return f.Create()
}

// Rename the file
func (f *File) Rename(targetName string) error {

	return os.Rename(f.FullPath, targetName)
}
