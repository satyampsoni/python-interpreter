
## Overview

This project implements a Python-like interpreter written in Go, demonstrating the fundamental concepts of language processing and execution.

## Concepts Behind Interpreters

### Interpreter Architecture

The interpreter follows a classic three-stage pipeline:

1. **Lexical Analysis (Lexing)**: Transforms source code text into a stream of tokens.
2. **Syntax Analysis (Parsing)**: Organizes tokens into an Abstract Syntax Tree (AST).
3. **Evaluation**: Executes the AST to produce program output.

### Tokenization and Lexical Analysis

Tokenization is the process of converting raw source code into a sequence of meaningful tokens. For example, the Python code:

```python
x = 5 + y
```

Is tokenized as:

| Token Type | Literal Value |
|------------|---------------|
| Identifier | `x`           |
| Operator   | `=`           |
| Number     | `5`           |
| Operator   | `+`           |
| Identifier | `y`           |

The lexer (`lexer.go`) handles this transformation by:

- Reading input character by character.
- Recognizing patterns like identifiers, numbers, and operators.
- Handling Python's significant whitespace (indentation).
- Producing a stream of tokens with their types and literal values.

### Parsing

The parser (`parser.go`) takes the token stream and builds an Abstract Syntax Tree (AST) representing the program structure. For a simple expression like `x = 5 + y`, the AST might look like:

```
Assignment
├── Identifier: x
└── Expression
    ├── Number: 5
    └── Identifier: y
```

The parser handles:

- Expressions (arithmetic, comparisons).
- Statements (assignments, function definitions, if/else).
- Block structure through indentation.
- Operator precedence and associativity.

### Evaluation

The evaluator (`evaluator.go`) traverses the AST and executes the program:

- It maintains an environment storing variables and their values.
- Expressions are evaluated to produce values.
- Statements are executed for their side effects.
- Control flow is managed (if/else conditionals).
- Function calls are dispatched.

## Project Structure

- `token/token.go`: Defines token types and structures.
- `lexer.go`: Implements lexical analysis.
- `parser.go`: Implements syntax analysis.
- `evaluator.go`: Implements evaluation.
- `environment.go`: Manages variable environments.
- `main.go`: REPL and entry point.

## Getting Started

### Prerequisites

- Go 1.16 or higher.

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/python-interpreter-go.git
```

Build the interpreter:

```bash
go build
```

Run the interactive REPL:

```bash
./python-interpreter-go
```

### Usage Examples

The interpreter supports basic Python-like syntax:

```python
x = 10
y = x + 5
print(y)
```

### AST Visualization Example

For the simple function:

```python
def add(a, b):
    return a + b
```

The AST would look like:

```
FunctionDefinition
├── Name: add
├── Parameters
│   ├── Identifier: a
│   └── Identifier: b
└── Body
    └── ReturnStatement
        └── Expression
            ├── Identifier: a
            └── Identifier: b
```

## Running Tests

Run tests using:

```bash
go test ./...
```

## Future Scope

The interpreter currently implements a subset of Python. Future enhancements could include:

### Advanced Language Features

- For loops and while loops.
- Lists, dictionaries, and other data structures.
- Classes and object-oriented programming.
- Exception handling (`try/except`).

### Standard Library

- File I/O operations.
- More built-in functions.
- Module system and imports.

### Performance Optimizations

- JIT compilation.
- Bytecode compilation.
- Optimized data structures.

### Developer Tools

- Improved error messages and debugging.
- Interactive debugger.
- Code completion in REPL.

### Compatibility

- Better Python compatibility.
- Support for external Python packages.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This interpreter was inspired by:

- Thorsten Ball's *"Writing an Interpreter in Go"*.
- Python's design and syntax.
- The rich history of programming language implementation.
