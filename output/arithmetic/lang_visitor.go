// Code generated from arithmetic/lang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by langParser.
type langVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by langParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by langParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
