package main

import (
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (int, error) {
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}
