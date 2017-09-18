// +build darwin

package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Printf("%#v\n", syscall.Kevent_t{})
}
