package action

import (
	"github.com/jakeslee/ikuai/action"
)

type HomepageType string

const (
	HomepageTypeSysStat  HomepageType = "sysstat"
	HomepageTypeACStatus HomepageType = "ac_status"
	HomepageTypeWANStat  HomepageType = "wan_stat"
	HomepageTypeWANSpeed HomepageType = "wan_speed"
)

func NewShowHomepageAction(typeNames ...HomepageType) *action.Action {
	return &action.Action{
		Action:   "show",
		FuncName: "homepage",
		Param: map[string]interface{}{
			"TYPE": Join(typeNames, ",", HomepageTypeSysStat),
		},
	}
}

type ShowHomepageResult struct {
	Status
	Results struct {
		SysStat  SysStat  `json:"sysstat,omitempty"`
		ACStatus ACStatus `json:"ac_status,omitempty"`
		WANStat  WANStat  `json:"wan_stat,omitempty"`
		Download int64    `json:"download,omitempty"`
		Upload   int64    `json:"upload,omitempty"`
	} `json:"results"`
}

type ACStatus struct {
	APCount  int `json:"ap_count"`
	APOnline int `json:"ap_online"`
}

type WANStat struct {
	Interface       string  `json:"interface"`
	ParentInterface string  `json:"parent_interface"`
	IPAddr          string  `json:"ip_addr"`
	Gateway         string  `json:"gateway"`
	Internet        string  `json:"internet"`
	Updatetime      string  `json:"updatetime"`
	AutoSwitch      string  `json:"auto_switch"`
	Result          string  `json:"result"`
	Errmsg          string  `json:"errmsg"`
	Comment         string  `json:"comment"`
	ISP             string  `json:"isp"`
	Rtt             string  `json:"rtt"`
	Total           int64   `json:"total"`
	Upload          float64 `json:"upload"`
	Download        float64 `json:"download"`
}
