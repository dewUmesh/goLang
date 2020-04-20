package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

/*
	This code create a new container in a different namespace
	Usage:
	go build container.go
	./container run /bin/sh
	or
	./container run <command> <string parameters>
*/

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what ???")

	}

}

func run() {
	fmt.Printf("%v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Run()
}
