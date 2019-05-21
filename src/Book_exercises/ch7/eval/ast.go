// abstract syntax tree
// as far as i got it
package myeval

// Expr represents an arithmetic expression.
type Expr interface {
	// Eval returns value of given Expr in Env enviroment.
	Eval(env Env) float64
	// Check reports about errors in passed Expr and add its Vars.
	Check(vars map[Var]bool) error
}

// Var determines variable, for example x.
type Var string

// literal represents numeric constant, for example 3.141.
type literal float64

// unary represents expression with unary operator, for example -x.
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binary represents expression with binary operator, for example x+y.
type binary struct {
	op   rune // '+', '-', '*' or '/'
	x, y Expr
}

// call represents expression of function call, for example sin(x).
type call struct {
	fn   string // one from "pow", "sin", "sqrt"
	args []Expr
}
