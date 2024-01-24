package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

// Start is the entry point for the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// Print the prompt
		fmt.Printf(PROMPT)

		// Read a line of input
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Get the line of input
		line := scanner.Text()

		// Create a lexer
		l := lexer.New(line)

		// Loop through the tokens
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// Print the token
			fmt.Printf("%+v\n", tok)
		}
	}
}
