package repl

import (
    "bufio"
    "fmt"
    "interpreter/evaluator"
    "interpreter/lexer"
    "interpreter/parser"
    "io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := evaluator.NewEnvironment()

    for {
        fmt.Fprintf(out, PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParserErrors(out, p.Errors())
            continue
        }

        evaluated := evaluator.Eval(program, env)
        if evaluated != nil {
            fmt.Fprintf(out, "%s\n", evaluated.Inspect())
        }
    }
}

func printParserErrors(out io.Writer, errors []string) {
    fmt.Fprintf(out, "Parser errors:\n")
    for _, msg := range errors {
        fmt.Fprintf(out, "\t%s\n", msg)
    }
}