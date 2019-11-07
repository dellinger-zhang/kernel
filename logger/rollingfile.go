package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"syscall"
	"errors"
)

const (
	checkInterval = 2 * time.Minute
)

var (
	filePipeline chan *RollingFile
)

type RollingFile struct {
	sync.Mutex
	fileName    string
	checkFlag   string
	fileHandler *os.File
	checkFun    func() string
}

func NewRollingFile(fileName string) *RollingFile {

	rf := &RollingFile{
		fileName: fileName,
		checkFun: func() string {
			return time.Now().Format("20060102")
		},
	}

	return rf
}

func (rf *RollingFile) SetCheckFun(fn func() string) {
	rf.checkFun = fn
}

func (rf *RollingFile) Open() error {

	rf.Lock()
	defer rf.Unlock()
	if rf.fileHandler != nil {
		rf.fileHandler.Close()
		rf.fileHandler = nil
	}
	rf.checkFlag = rf.checkFun()
	fileName := fmt.Sprintf("%s.%s", rf.fileName, rf.checkFlag)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Errorf("open file error:%s", fileName)
		return err
	}
	rf.fileHandler = f

	filePipeline <- rf
	return nil
}

func (rf *RollingFile) Write(p []byte) (n int, err error) {
	if rf.fileHandler != nil {
		rf.Lock()
		defer rf.Unlock()
		return rf.fileHandler.Write(p)
	}
	return 0, os.ErrClosed
}

func (rf *RollingFile) RedirectSTDOutput() error {
	if rf.fileHandler != nil {
		if err := syscall.Dup2(int(rf.fileHandler.Fd()), 1); err != nil {
			return err
		}
		if err := syscall.Dup2(int(rf.fileHandler.Fd()), 2); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("RedirectSTDOutput:Open file first")
	}
}

func _startTimer() {
	tick := time.Tick(checkInterval)
	go func() {
		fileList := make(map[*RollingFile]interface{})
		for {
			select {
			case f := <-filePipeline:
				fileList[f] = new(interface{})
			case <-tick:
				for k, _ := range fileList {
					newName := k.checkFun()
					if !strings.EqualFold(newName, k.fileName) {
						k.checkFlag = newName
						delete(fileList, k)
						go k.Open()
					}
				}
			}
		}
	}()
}

func init() {
	filePipeline = make(chan *RollingFile)
	_startTimer()
}
