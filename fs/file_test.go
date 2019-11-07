package fs

import "testing"

func TestIsDir(t *testing.T) {

	homeDir, err := Home()
	if err != nil {
		t.Error("homeDir fail")
		t.FailNow()
	}

	f := NewFile(homeDir)
	if !f.Exists() {
		t.Error("dir exists fail")
		t.Fail()
	}

	if !IsDir(f.FullPath) {
		t.Error("dir judged fail")
		t.Fail()
	}
}
