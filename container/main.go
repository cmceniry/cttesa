package main

import (
        "os"
        "os/exec"
        "syscall"
)

func main() {
        cmd := exec.Command("/bin/bash")
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        cmd...

        err := cmd.Run()
        if err != nil {
                panic(err)
        }
}
