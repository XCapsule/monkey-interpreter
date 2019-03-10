package main

import (
	"fmt"
	"github.com/XCapsule/monkey-interpreter/repl"
	"log"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("err when start:%v", err)
		return
	}
	fmt.Printf("Welcome Monkey Interpreter World, MR/MS %s!\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
