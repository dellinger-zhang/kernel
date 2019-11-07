package fs

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// ROOT indicates the root of current app.
	ROOT = "/"
)

const (
	ReadMode      = 0444
	WriteMode     = 0222
	ReadWriteMode = 0666
	CommonMode    = 0766
	FullMode      = 0777
)

// IsDir checks input whether directory or not.
func IsDir(dir string) bool {

	fhandler, err := os.Stat(dir)
	if !(err == nil || os.IsExist(err)) {
		return false
	}
	return fhandler.IsDir()
}

// IsFile checks input whether file or not.
func IsFile(filename string) bool {
	fhandler, err := os.Stat(filename)
	if !(err == nil || os.IsExist(err)) {
		return false
	} else if fhandler.IsDir() {
		return false
	}
	return true
}

// GetFileSize returns bytes of file.
func GetFileSize(filename string) (bool, int64) {

	if !IsFile(filename) {
		return false, 0
	}
	fhandler, _ := os.Stat(filename)
	return true, fhandler.Size()
}

// MakeDirP equals to mkdir -p
func MakeDirP(dir string) error {

	return os.MkdirAll(dir, os.ModePerm)
}

// CurrentDir returns the folder of program
func currentDir() string {
	ret, err := os.Getwd()
	if err != nil {
		s, err := exec.LookPath(os.Args[0])
		if err != nil {
			ret = ""
			return ret
		}

		i := strings.LastIndex(s, string(os.PathSeparator))
		ret = string(s[0:i])
	}
	return ret
}

// FileExists checks file exist or not
func FileExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return err == nil || os.IsExist(err)
}

// CombinPath combins all paramters to whole path.
func CombinPath(elements ...string) string {

	return strings.Join(elements, string(os.PathSeparator))
}

// RemoveFile to delete a file
func RemoveFile(fullPath string) bool {

	return os.Remove(fullPath) == nil
}

// CopyFile to copy src file to dest file.
func CopyFile(srcName, destName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		return -1, err
	}
	defer src.Close()

	dst, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return -1, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func SymLink(srcName, destName string) error {

	destDir := filepath.Dir(destName)
	if err := MakeDirP(destDir); err != nil {
		return err
	}

	return os.Symlink(srcName, destName)
}

func WalkDir(dir string, fn func(*File)) error {

	if !FileExists(dir) {
		return os.ErrNotExist
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err == nil {
			fn(NewFile(path))
		}
		return nil
	})
	return nil
}

func LockFile(file string) bool {
	f := NewFile(file)
	if f.Exists() {
		return false
	}
	f.Create()
	return true
}

func UnlockFile(file string) bool {
	return NewFile(file).Delete()
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func init() {

	ROOT = currentDir()
}
