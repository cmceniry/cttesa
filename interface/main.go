package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func saveFile(filename string, data io.Reader) error {
	buf := make([]byte, 1024)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	nr, err := data.Read(buf)
	if err != nil {
		return err
	}
	nw, err := file.Write(buf[:nr])
	if err != nil {
		return err
	}
	if nw != nr {
		return fmt.Errorf("Bytes Written don't match Bytes Read: %s", filename)
	}
	return nil
}

func main() {
	a := bytes.NewBuffer([]byte("file1"))
	err := saveFile("file1\n", a)
	if err != nil {
		panic(err)
	}
	b := strings.NewReader("file2")
	err = saveFile("file2\n", b)
	if err != nil {
		panic(err)
	}
}
