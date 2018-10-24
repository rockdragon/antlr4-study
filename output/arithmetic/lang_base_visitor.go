// Code generated from arithmetic/lang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaselangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaselangVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
