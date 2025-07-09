// The comments in this file are inspired by Harvey Specter - my fav character from suits 

package evaluator

import (
    "fmt"
    "interpreter/parser"
    "strconv"
)

// Look, we need types to represent different kinds of data. This isn't a democracy.
type ObjectType string

const (
    INTEGER_OBJ  = "INTEGER"
    STRING_OBJ   = "STRING"
    BOOLEAN_OBJ  = "BOOLEAN"
    NULL_OBJ     = "NULL"
    ERROR_OBJ    = "ERROR"
    BUILTIN_OBJ  = "BUILTIN"
)

// Everything's an Object. Deal with it.
type Object interface {
    Type() ObjectType
    Inspect() string
}

// NULL - because sometimes nothing is better than something
var NULL = &NullObject{}

type NullObject struct{}

func (n *NullObject) Type() ObjectType { return NULL_OBJ }
func (n *NullObject) Inspect() string  { return "null" }

// Integers. They're simple. I like simple.
type Integer struct {
    Value string
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return i.Value }

// Strings. For when numbers aren't enough.
type String struct {
    Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

func (s *String) QuotedValue() string {
    return fmt.Sprintf("\"%s\"", s.Value)
}

// Booleans. True or false. Like my cases - I only take the ones I'll win.
type Boolean struct {
    Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  {
    if b.Value {
        return "true"
    }
    return "false"
}

// Errors. They happen. I fix them.
type Error struct {
    Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Functions that are built-in. Like my charm.
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
    Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// Built-ins. They're not up for negotiation.
var builtins = map[string]*Builtin{
    "print": {
        Fn: func(args ...Object) Object {
            for _, arg := range args {
                fmt.Println(arg.Inspect())
            }
            return NULL
        },
    },
    // When I need more, I'll add them. And they'll be spectacular.
}

// Eval - the closer. It handles every case and never loses.
func Eval(node interface{}, env *Environment) Object {
    switch node := node.(type) {
    case *parser.Program:
        return evalProgram(node, env)
        
    case *parser.ExpressionStatement:
        return Eval(node.Expression, env)
        
    case *parser.IntegerLiteral:
        return &Integer{Value: node.Value}
        
    case *parser.StringLiteral:
        return &String{Value: node.Value}
        
    case *parser.InfixExpression:
        left := Eval(node.Left, env)
        if isError(left) {
            return left
        }
        right := Eval(node.Right, env)
        if isError(right) {
            return right
        }
        return evalInfixExpression(node.Operator, left, right)
        
    case *parser.Identifier:
        return evalIdentifier(node, env)
        
    case *parser.AssignmentStatement:
        val := Eval(node.Value, env)
        if isError(val) {
            return val
        }
        return env.Set(node.Name.Value, val)

    case *parser.CallExpression:
        function := Eval(node.Function, env)
        if isError(function) {
            return function
        }
        args := evalExpressions(node.Arguments, env)
        if len(args) == 1 && isError(args[0]) {
            return args[0]
        }
        return applyFunction(function, args)
    default:
        return newError("unknown node type")
    }
}

// Programs are just a series of statements. Like my winning streak.
func evalProgram(program *parser.Program, env *Environment) Object {
    var result Object = NULL
    
    for _, statement := range program.Statements {
        result = Eval(statement, env)
        
        if isError(result) {
            return result
        }
    }
    
    return result
}

// Look up identifiers. I always know who I'm dealing with.
func evalIdentifier(node *parser.Identifier, env *Environment) Object {
    if builtin, ok := builtins[node.Value]; ok {
        return builtin
    }
    
    if val, ok := env.Get(node.Value); ok {
        return val
    }

    return newError("identifier not found: %s", node.Value)
}

// Evaluate expressions. I do this with witnesses all the time.
func evalExpressions(exps []parser.Expression, env *Environment) []Object {
    var result []Object
    
    for _, e := range exps {
        evaluated := Eval(e, env)
        if isError(evaluated) {
            return []Object{evaluated}
        }
        result = append(result, evaluated)
    }
    
    return result
}

// Apply functions. Like applying pressure to get what I want.
func applyFunction(fn Object, args []Object) Object {
    switch fn := fn.(type) {
    case *Builtin:
        return fn.Fn(args...)
    default:
        return newError("not a function: %s", fn.Type())
    }
}

// Is it an error? I need to know before I proceed.
func isError(obj Object) bool {
    if obj != nil && obj.Type() == ERROR_OBJ {
        return true
    }
    return false
}

// Evaluate operations between values. Math never lies.
func evalInfixExpression(operator string, left Object, right Object) Object {
    switch {
    case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
        return evalIntegerInfixExpression(operator, left.(*Integer), right.(*Integer))
    case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
        return evalStringInfixExpression(operator, left.(*String), right.(*String))
    default:
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
    }
}

// Integer operations. Clean, precise, and final. Like my arguments in court.
func evalIntegerInfixExpression(operator string, left *Integer, right *Integer) Object {
    leftVal := left.Value
    rightVal := right.Value

    leftInt, leftErr := strconv.Atoi(leftVal)
    rightInt, rightErr := strconv.Atoi(rightVal)
    
    if leftErr != nil || rightErr != nil {
        return newError("invalid integer operation: %s %s %s", leftVal, operator, rightVal)
    }

    var result int
    switch operator {
    case "+":
        result = leftInt + rightInt
    case "-":
        result = leftInt - rightInt
    case "*":
        result = leftInt * rightInt
    case "/":
        if rightInt == 0 {
            return newError("division by zero")
        }
        result = leftInt / rightInt
    default:
        return newError("unknown operator: %s", operator)
    }
    
    return &Integer{Value: strconv.Itoa(result)}
}

// String operations. Sometimes words are more powerful than numbers.
func evalStringInfixExpression(operator string, left *String, right *String) Object {
    if operator != "+" {
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
    }
    
    return &String{Value: left.Value + right.Value}
}

// Create errors with style and precision.
func newError(format string, a ...interface{}) *Error {
    return &Error{Message: fmt.Sprintf(format, a...)}
}