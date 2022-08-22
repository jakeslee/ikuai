package action

// NewStreamIPPortSwitchAction 上下线分流
func NewStreamIPPortSwitchAction(id string, state SwitchState) *Action {
	return NewSwitchAction(&Action{
		FuncName: "stream_ipport",
		Param: map[string]interface{}{
			"id": id,
		},
	}, state)
}
