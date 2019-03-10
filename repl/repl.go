package repl

import (
	"bufio"
	"fmt"
	"github.com/XCapsule/monkey-interpreter/lexer"
	"github.com/XCapsule/monkey-interpreter/token"
	"io"
	"log"
)

const PROMPT = "^_^ >>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			if t.Type == token.ILLEGAL {
				log.Fatal("ILLEGAL Input!")
				return
			}
			fmt.Printf("%+v\n", t)
		}
	}
}
