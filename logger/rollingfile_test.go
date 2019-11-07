package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRollingFile(t *testing.T) {

	rf := NewRollingFile("/Users/dellinger/work/test")
	rf.SetCheckFun(func() string {
		return time.Now().Format("20160102.150405")
	})

	if err := rf.Open(); err != nil {
		t.Error(err)
		t.Fail()
	}

	for i := 0; i < 200; i++ {
		rf.Write([]byte(fmt.Sprintf("line %d: hello bibi\n", i)))
		time.Sleep(100 * time.Millisecond)
	}
}