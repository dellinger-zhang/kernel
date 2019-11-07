package fs

import (
	"git.tutorabc.com/tmc3/corewars/kernel.git/util"
	"path"
	"strings"
	"testing"
	"time"
)

func TestWalkDir(t *testing.T) {

	dir := "/Users/dellinger/Documents"
	WalkDir(dir, func(file *File) {

		outdir := "/Users/dellinger/temp/output"

		if util.IsEmptyStr(file.FileRawName) || file.FileRawName[0] == '.' || IsDir(file.FullPath) {
			return
		}

		relPath := strings.Replace(file.FullPath, dir, "", -1)
		if file.FileExtName == ".json" {
			CopyFile(file.FullPath, path.Join(outdir, relPath))
		} else {
			SymLink(file.FullPath, path.Join(outdir, relPath))
		}
	})

	time.Sleep(5 * time.Second)
}
