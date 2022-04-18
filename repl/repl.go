package repl

import (
	"bufio"
	"fmt"
	"github.com/tatsuki1112/writing-an-interpreter-in-go/writing-an-interpreter-in-go/lexer"
	"github.com/tatsuki1112/writing-an-interpreter-in-go/writing-an-interpreter-in-go/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken() ; tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}