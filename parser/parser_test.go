package parser

import (
    "interpreter/lexer"
    "testing"
)

func TestIntegerLiteralExpression(t *testing.T) {
    input := "5"

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
    }

    literal, ok := stmt.Expression.(*IntegerLiteral)
    if !ok {
        t.Fatalf("stmt.Expression is not IntegerLiteral. got=%T", stmt.Expression)
    }

    if literal.Value != "5" {
        t.Errorf("literal.Value not %s. got=%s", "5", literal.Value)
    }
}

func TestInfixExpression(t *testing.T) {
    input := "5 + 3 * 2"

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
    }

    exp, ok := stmt.Expression.(*InfixExpression)
    if !ok {
        t.Fatalf("stmt.Expression is not InfixExpression. got=%T", stmt.Expression)
    }

    if exp.Operator != "+" {
        t.Errorf("exp.Operator is not '+'. got=%s", exp.Operator)
    }

    left, ok := exp.Left.(*IntegerLiteral)
    if !ok {
        t.Fatalf("exp.Left is not IntegerLiteral. got=%T", exp.Left)
    }

    if left.Value != "5" {
        t.Errorf("left.Value not %s. got=%s", "5", left.Value)
    }

    right, ok := exp.Right.(*InfixExpression)
    if !ok {
        t.Fatalf("exp.Right is not InfixExpression. got=%T", exp.Right)
    }

    if right.Operator != "*" {
        t.Errorf("right.Operator is not '*'. got=%s", right.Operator)
    }

    rightLeft, ok := right.Left.(*IntegerLiteral)
    if !ok {
        t.Fatalf("right.Left is not IntegerLiteral. got=%T", right.Left)
    }

    if rightLeft.Value != "3" {
        t.Errorf("rightLeft.Value not %s. got=%s", "3", rightLeft.Value)
    }

    rightRight, ok := right.Right.(*IntegerLiteral)
    if !ok {
        t.Fatalf("right.Right is not IntegerLiteral. got=%T", right.Right)
    }

    if rightRight.Value != "2" {
        t.Errorf("rightRight.Value not %s. got=%s", "2", rightRight.Value)
    }
}

func TestGroupedExpression(t *testing.T) {
    input := "(5 + 3) * 2"

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
    }

    exp, ok := stmt.Expression.(*InfixExpression)
    if !ok {
        t.Fatalf("stmt.Expression is not InfixExpression. got=%T", stmt.Expression)
    }

    if exp.Operator != "*" {
        t.Errorf("exp.Operator is not '*'. got=%s", exp.Operator)
    }

    left, ok := exp.Left.(*InfixExpression)
    if !ok {
        t.Fatalf("exp.Left is not InfixExpression. got=%T", exp.Left)
    }

    if left.Operator != "+" {
        t.Errorf("left.Operator is not '+'. got=%s", left.Operator)
    }

    leftLeft, ok := left.Left.(*IntegerLiteral)
    if !ok {
        t.Fatalf("left.Left is not IntegerLiteral. got=%T", left.Left)
    }

    if leftLeft.Value != "5" {
        t.Errorf("leftLeft.Value not %s. got=%s", "5", leftLeft.Value)
    }

    leftRight, ok := left.Right.(*IntegerLiteral)
    if !ok {
        t.Fatalf("left.Right is not IntegerLiteral. got=%T", left.Right)
    }

    if leftRight.Value != "3" {
        t.Errorf("leftRight.Value not %s. got=%s", "3", leftRight.Value)
    }

    right, ok := exp.Right.(*IntegerLiteral)
    if !ok {
        t.Fatalf("exp.Right is not IntegerLiteral. got=%T", exp.Right)
    }

    if right.Value != "2" {
        t.Errorf("right.Value not %s. got=%s", "2", right.Value)
    }
}

func TestParseFunctionDefinition(t *testing.T) {
    input := `def myFunc(x, y):
    return x + y`

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
    }

    funcDef, ok := program.Statements[0].(*FunctionDefinition)
    if !ok {
        t.Fatalf("program.Statements[0] is not FunctionDefinition. got=%T", program.Statements[0])
    }

    if funcDef.Name != "myFunc" {
        t.Errorf("funcDef.Name not %s. got=%s", "myFunc", funcDef.Name)
    }

    if len(funcDef.Parameters) != 2 {
        t.Fatalf("funcDef.Parameters does not contain 2 parameters. got=%d", len(funcDef.Parameters))
    }

    if funcDef.Parameters[0] != "x" || funcDef.Parameters[1] != "y" {
        t.Errorf("funcDef.Parameters not [x, y]. got=%v", funcDef.Parameters)
    }

    if len(funcDef.Body) != 1 {
        t.Fatalf("funcDef.Body does not contain 1 statement. got=%d", len(funcDef.Body))
    }

    returnStmt, ok := funcDef.Body[0].(*ReturnStatement)
    if !ok {
        t.Fatalf("funcDef.Body[0] is not ReturnStatement. got=%T", funcDef.Body[0])
    }

    infixExp, ok := returnStmt.Value.(*InfixExpression)
    if !ok {
        t.Fatalf("returnStmt.Value is not InfixExpression. got=%T", returnStmt.Value)
    }

    if infixExp.Operator != "+" {
        t.Errorf("infixExp.Operator not '+'. got=%s", infixExp.Operator)
    }
}

func TestParseReturnStatement(t *testing.T) {
    input := `return 42`

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
    }

    returnStmt, ok := program.Statements[0].(*ReturnStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not ReturnStatement. got=%T", program.Statements[0])
    }

    literal, ok := returnStmt.Value.(*IntegerLiteral)
    if !ok {
        t.Fatalf("returnStmt.Value is not IntegerLiteral. got=%T", returnStmt.Value)
    }

    if literal.Value != "42" {
        t.Errorf("literal.Value not %s. got=%s", "42", literal.Value)
    }
}

func TestParseAssignmentStatement(t *testing.T) {
    input := `x = 42`

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
    }

    assignStmt, ok := program.Statements[0].(*AssignmentStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not AssignmentStatement. got=%T", program.Statements[0])
    }

    if assignStmt.Name.Value != "x" {
        t.Errorf("assignStmt.Name.Value not %s. got=%s", "x", assignStmt.Name.Value)
    }

    literal, ok := assignStmt.Value.(*IntegerLiteral)
    if !ok {
        t.Fatalf("assignStmt.Value is not IntegerLiteral. got=%T", assignStmt.Value)
    }

    if literal.Value != "42" {
        t.Errorf("literal.Value not %s. got=%s", "42", literal.Value)
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %s", msg)
    }
    t.FailNow()
}