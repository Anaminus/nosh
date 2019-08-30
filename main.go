package main

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	switch len(os.Args) {
	case 0:
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
	}
	cmd.Start()
}
