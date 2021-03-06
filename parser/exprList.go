package parser

import (
	"github.com/macbinn/hacklang/ast"
	"github.com/macbinn/hacklang/token"
)

type exprListHandler struct {

}

func (exprListHandler) Parse(tokens []*token.Token) (ast.Node, int, error) {
	node := new(ast.ExprList)
	pos := 0
	for pos < len(tokens) {
		expr, i, err := ParseGreedy(tokens[pos:],
			"expr",
		)
		if err != nil {
			break
		}
		node.Nodes = append(node.Nodes, expr)
		pos += i
	}
	if len(node.Nodes) == 0 {
		return nil, 0, ErrSyntaxError
	}
	return node, pos, nil
}

func init() {
	Register("exprList", exprListHandler{})
}
