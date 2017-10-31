package main

import (
	"fmt"
	"os"
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

func main() {
	ino, err := getInode(".")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inode: %d\n", ino)
}
