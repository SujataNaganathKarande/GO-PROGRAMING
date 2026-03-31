package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

func main() {
	var rw ReadWriter
	rw = &File{}

	rw.Write("Go interfaces are powerful!")
	fmt.Println("File content:", rw.Read())
}
