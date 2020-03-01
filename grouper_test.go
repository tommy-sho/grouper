package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type fakeFileInfo string

func (fi fakeFileInfo) Name() string    { return string(fi) }
func (fakeFileInfo) Sys() interface{}   { return nil }
func (fakeFileInfo) ModTime() time.Time { return time.Time{} }
func (fakeFileInfo) IsDir() bool        { return false }
func (fakeFileInfo) Size() int64        { return 0 }
func (fakeFileInfo) Mode() os.FileMode  { return 0644 }

func Test_isGoFile(t *testing.T) {
	testGoFile := fakeFileInfo("test.go")
	if !isGoFile(testGoFile) {
		t.Errorf(fmt.Sprintf("must determine %s as Go file", testGoFile))
	}

	testNoGoFile := fakeFileInfo("test.text")
	if isGoFile(testNoGoFile) {
		t.Errorf(fmt.Sprintf("must determine %s as non Go file", testNoGoFile))
	}
}
