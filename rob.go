package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("Ver bana en az 1 arguman")
	}
	//parsePath(os.Getenv("PATH"))
	cmd := exec.Command(args[1], args[2:]...)

	//cmd.Path = os.Getenv("PATH")

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	errHandler.HandlerExit(err)
	os.Exit(0)

}
