package action

import "github.com/jakeslee/ikuai/action"

type RouteGroup struct {
	Tagname    string       `json:"tagname"`
	GroupID    string       `json:"group_id"`
	RefCount   int64        `json:"ref_count"`
	GroupValue []GroupValue `json:"group_value"`
	Type       int64        `json:"type"`
	ID         int64        `json:"id"`
	GroupName  string       `json:"group_name"`
}

type GroupValue struct {
	IP      string `json:"ip"`
	Comment string `json:"comment"`
}

type ShowRouteGroupResult struct {
	Status
	Results struct {
		Data  []RouteGroup `json:"data"`
		Total int64        `json:"total"`
	} `json:"results"`
}

func NewShowRouteObjectIPAction() *action.Action {
	return &action.Action{
		Action:   "show",
		FuncName: "route_object_ip",
		Param: map[string]interface{}{
			"TYPE":  "data,total",
			"limit": "0,10000",
		},
	}
}

func NewShowRouteObjectIPv6Action() *action.Action {
	return &action.Action{
		Action:   "show",
		FuncName: "route_object_ip6",
		Param: map[string]interface{}{
			"TYPE":  "data,total",
			"limit": "0,10000",
		},
	}
}
