package action

type IfaceCheck struct {
	ID              int    `json:"id"`
	Interface       string `json:"interface"`
	ParentInterface string `json:"parent_interface"`
	IPAddr          string `json:"ip_addr"`
	Gateway         string `json:"gateway"`
	Internet        string `json:"internet"`
	Updatetime      string `json:"updatetime"`
	AutoSwitch      string `json:"auto_switch"`
	Result          string `json:"result"`
	Errmsg          string `json:"errmsg"`
	Comment         string `json:"comment"`
}

type IfaceStream struct {
	Interface   string `json:"interface"`
	Comment     string `json:"comment"`
	IPAddr      string `json:"ip_addr"`
	ConnectNum  string `json:"connect_num"`
	Upload      int    `json:"upload"`
	Download    int    `json:"download"`
	TotalUp     int64  `json:"total_up"`
	TotalDown   int64  `json:"total_down"`
	Updropped   int    `json:"updropped"`
	Downdropped int    `json:"downdropped"`
	Uppacked    int    `json:"uppacked"`
	Downpacked  int    `json:"downpacked"`
}

func NewMonitorInterfaceAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "monitor_iface",
		Param: map[string]interface{}{
			"TYPE": "iface_check,iface_stream",
		},
	}
}

type ShowMonitorInterfaceResult struct {
	Result
	Data struct {
		IfaceCheck  []IfaceCheck  `json:"iface_check"`
		IfaceStream []IfaceStream `json:"iface_stream"`
	} `json:"data"`
}
