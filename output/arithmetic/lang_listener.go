// Code generated from arithmetic/lang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// langListener is a complete listener for a parse tree produced by langParser.
type langListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
