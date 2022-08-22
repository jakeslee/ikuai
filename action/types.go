package action

import "strings"

type Action struct {
	Action   string                 `json:"action"`
	FuncName string                 `json:"func_name"`
	Param    map[string]interface{} `json:"param,omitempty"`
}

type SwitchState string

const (
	SwitchStateUp   SwitchState = "up"
	SwitchStateDown SwitchState = "down"
)

func NewSwitchAction(a *Action, state SwitchState) *Action {
	return &Action{
		Action:   string(state),
		FuncName: a.FuncName,
		Param:    a.Param,
	}
}

// Result 操作结果
type Result struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
}

type DataResult struct {
	Total int                      `json:"total"`
	Data  []map[string]interface{} `json:"data"`
}

// IPGroup IP 组信息
type IPGroup struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
	AddrPool  string `json:"addr_pool"`
	Comment   string `json:"comment"`
}

func (i *IPGroup) AddIPs(ips []string) {
	i.AddrPool = strings.Join(ips, ",")
}

func (i *IPGroup) AddComments(comments []string) {
	i.Comment = strings.Join(comments, ",")
}
