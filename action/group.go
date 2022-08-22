package action

// NewIPGroupEditAction 修改 IP 组
func NewIPGroupEditAction(ipGroup IPGroup) *Action {
	return &Action{
		Action:   "edit",
		FuncName: "ipgroup",
		Param: map[string]interface{}{
			"id":         ipGroup.Id,
			"group_name": ipGroup.GroupName,
			"addr_pool":  ipGroup.AddrPool,
			"comment":    ipGroup.Comment,
		},
	}
}

// ShowIPGroupResult IP 组结果
type ShowIPGroupResult struct {
	Result

	Data struct {
		Data  []IPGroup `json:"data"`
		Total int       `json:"total"`
	} `json:"Data"`
}

// NewIPGroupShowAction 查看 IP 组列表
func NewIPGroupShowAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "ipgroup",
	}
}
