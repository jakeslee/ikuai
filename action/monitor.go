package action

type LanDeviceInfo struct {
	Webid        int    `json:"webid"`
	Downrate     string `json:"downrate"`
	TotalUp      int    `json:"total_up"`
	TotalDown    int    `json:"total_down"`
	Mac          string `json:"mac"`
	Uprate       string `json:"uprate"`
	Uptime       string `json:"uptime"`
	Hostname     string `json:"hostname"`
	DtalkName    string `json:"dtalk_name"`
	Signal       string `json:"signal"`
	AcGid        int    `json:"ac_gid"`
	Frequencies  string `json:"frequencies"`
	Bssid        string `json:"bssid"`
	Reject       int    `json:"reject"`
	ClientType   string `json:"client_type"`
	IPAddrInt    int    `json:"ip_addr_int"`
	ConnectNum   int    `json:"connect_num"`
	Upload       int    `json:"upload"`
	Download     int    `json:"download"`
	AuthType     int    `json:"auth_type"`
	IPAddr       string `json:"ip_addr"`
	ClientDevice string `json:"client_device"`
	Timestamp    int    `json:"timestamp"`
	Comment      string `json:"comment"`
	Username     string `json:"username"`
	Ppptype      string `json:"ppptype"`
	Apmac        string `json:"apmac"`
	Apname       string `json:"apname"`
	Ssid         string `json:"ssid"`
	ID           int    `json:"id"`
}

func NewMonitorLanIpAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "monitor_lanip",
		Param: map[string]interface{}{
			"TYPE": "data,total",
		},
	}
}

type ShowMonitorResult struct {
	Result
	Data struct {
		DataResult
		Data []LanDeviceInfo `json:"data"`
	} `json:"data"`
}

func NewShowSysStatAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "homepage",
		Param: map[string]interface{}{
			"TYPE": "sysstat",
		},
	}
}

type Verinfo struct {
	Modelname    string `json:"modelname"`
	Verstring    string `json:"verstring"`
	Version      string `json:"version"`
	BuildDate    int64  `json:"build_date"`
	Arch         string `json:"arch"`
	Sysbit       string `json:"sysbit"`
	Verflags     string `json:"verflags"`
	IsEnterprise int    `json:"is_enterprise"`
	SupportI18N  int    `json:"support_i18n"`
	SupportLcd   int    `json:"support_lcd"`
}

type MemoryStat struct {
	Total     int64  `json:"total"`
	Available int64  `json:"available"`
	Free      int64  `json:"free"`
	Cached    int64  `json:"cached"`
	Buffers   int64  `json:"buffers"`
	Used      string `json:"used"`
}

type StreamStat struct {
	ConnectNum int   `json:"connect_num"`
	Upload     int   `json:"upload"`
	Download   int   `json:"download"`
	TotalUp    int64 `json:"total_up"`
	TotalDown  int64 `json:"total_down"`
}

type SysStat struct {
	Hostname   string `json:"hostname"`
	Gwid       string `json:"gwid"`
	OnlineUser struct {
		Count int `json:"count"`
	} `json:"online_user"`
	Uptime  int        `json:"uptime"`
	Cpu     []string   `json:"cpu"`
	Freq    []string   `json:"freq"`
	Cputemp []int      `json:"cputemp"`
	Verinfo Verinfo    `json:"verinfo"`
	Memory  MemoryStat `json:"memory"`
	Stream  StreamStat `json:"stream"`
}

type ShowSysStatResult struct {
	Result
	Data struct {
		SysStat SysStat `json:"sysstat"`
	} `json:"data"`
}
