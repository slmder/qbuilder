package parts

import (
	"fmt"
	"strings"

	expression2 "github.com/slmder/qbuilder/parts/expression"
)

type OrderBy struct {
	OrderBy []expression2.SortExpression
}

func (p OrderBy) String() string {
	if len(p.OrderBy) > 0 {
		return joinOrderByExpressions(p.OrderBy)
	}
	return ""
}

func (p *OrderBy) Set(expr string, direction string) {
	p.OrderBy = []expression2.SortExpression{{Expression: expr, Direction: direction}}
}

func (p *OrderBy) Reset() {
	p.OrderBy = []expression2.SortExpression{}
}

func (p *OrderBy) Add(expr string, direction string) {
	p.OrderBy = append(p.OrderBy, expression2.SortExpression{Expression: expr, Direction: direction})
}

func joinOrderByExpressions(expressions []expression2.SortExpression) string {
	var res []string
	for _, expr := range expressions {
		res = append(res, expr.String())
	}
	return fmt.Sprintf("ORDER BY %s", strings.Join(res, ", "))
}
