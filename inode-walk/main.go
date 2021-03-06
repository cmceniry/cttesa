package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func getInode(path string) (uint64, error) {
	s, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	stat, ok := s.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, fmt.Errorf("Wrong type for stat")
	}
	return stat.Ino, nil
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return filepath.SkipDir
	}
	if info.IsDir() {
		return nil
	}
	ino, err := getInode(path)
	if err != nil {
		return err
	}
	fmt.Printf("%s = %d\n", path, ino)
	return nil
}

func main() {
	filepath.Walk(".", walkFunc)
}
