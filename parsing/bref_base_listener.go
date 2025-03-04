// Code generated from bref.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // bref
import "github.com/antlr4-go/antlr/v4"

// BasebrefListener is a complete listener for a parse tree produced by brefParser.
type BasebrefListener struct{}

var _ brefListener = &BasebrefListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebrefListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebrefListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebrefListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebrefListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasebrefListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasebrefListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasebrefListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasebrefListener) ExitExpr(ctx *ExprContext) {}
