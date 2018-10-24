// Code generated from arithmetic/lang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselangListener is a complete listener for a parse tree produced by langParser.
type BaselangListener struct{}

var _ langListener = &BaselangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaselangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaselangListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaselangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaselangListener) ExitExpr(ctx *ExprContext) {}
