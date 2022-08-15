package internal

import "strings"

type ConditionsGroup struct {
	condition     strings.Builder
	conditionArgs []any
}

func NewConditionsGroup(condition string, args ...any) *ConditionsGroup {
	group := &ConditionsGroup{conditionArgs: args}
	group.condition.WriteString(condition)
	return group
}

func (g *ConditionsGroup) AddANDCondition(condition string, args ...any) {
	g.addCondition("AND", condition, args...)
}

func (g *ConditionsGroup) AddORCondition(condition string, args ...any) {
	g.addCondition("OR", condition, args...)
}

func (g *ConditionsGroup) addCondition(conditionType, condition string, args ...any) {
	g.condition.WriteRune(' ')
	g.condition.WriteString(conditionType)
	g.condition.WriteRune(' ')
	g.condition.WriteString(condition)
	g.conditionArgs = append(g.conditionArgs, args...)
}

func (g *ConditionsGroup) GetGORMConds() []any {
	return append([]any{g.condition.String()}, g.conditionArgs...)
}
