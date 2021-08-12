package qbuilder

import (
	parts2 "github.com/sergeyKhashimov/qbuilder/parts"
	"strings"
)

type UpdateBuilder struct {
	builder
	update parts2.Update
	set    parts2.Set
	from   parts2.From
	where  parts2.Where
}

func (u *UpdateBuilder) Update(rel string) *UpdateBuilder {
	u.update = parts2.Update{Relation: rel}
	return u
}

func (u *UpdateBuilder) From(rel string) *UpdateBuilder {
	u.from = parts2.From{Relation: rel}
	return u
}

func (u *UpdateBuilder) Set(column string, value string) *UpdateBuilder {
	u.set.Add(column, value)
	return u
}

func (u *UpdateBuilder) Where(expr... string) *UpdateBuilder {
	u.where.Reset()
	for _, e := range expr {
		u.where.Add(e)
	}
	return u
}

func (u *UpdateBuilder) AndWhere(expr string) *UpdateBuilder {
	u.where.Add(expr)
	return u
}

func (u *UpdateBuilder) SetParameter(name string, value interface{}) *UpdateBuilder {
	u.parameters.Set(name, value)
	return u
}

func (u *UpdateBuilder) SetParameters(params map[string]interface{}) *UpdateBuilder {
	for name, value := range params {
		u.parameters.Set(name, value)
	}
	return u
}

func (u *UpdateBuilder) RemoveParameter(name string) *UpdateBuilder {
	u.parameters.Remove(name)
	return u
}

func (u *UpdateBuilder) ToSQL() string {
	expressions := []string{
		u.update.String(),
		u.set.String(),
		u.from.String(),
		u.where.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}