// +build linux

package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Printf("%#v\n", syscall.EpollEvent{})
}
