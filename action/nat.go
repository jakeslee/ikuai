package action

// NewNATRuleSwitchAction 切换 NAT 规则
func NewNATRuleSwitchAction(id string, state SwitchState) *Action {
	return NewSwitchAction(&Action{
		FuncName: "nat_rule",
		Param: map[string]interface{}{
			"id": id,
		},
	}, state)
}
