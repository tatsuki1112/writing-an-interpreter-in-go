package main

import (
	"fmt"
	"github.com/tatsuki1112/writing-an-interpreter-in-go/writing-an-interpreter-in-go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is the MOnkey Programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
