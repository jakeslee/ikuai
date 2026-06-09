package action

import "github.com/jakeslee/ikuai/action"

type LanDeviceInfo struct {
	Bssid        string  `json:"bssid"`
	Uprate       string  `json:"uprate"`
	Downrate     string  `json:"downrate"`
	Enc          string  `json:"enc"`
	Channel      string  `json:"channel"`
	ID           int64   `json:"id"`
	MAC          string  `json:"mac"`
	Download     int64   `json:"download"`
	IPAddrInt    int64   `json:"ip_addr_int"`
	ConnectNum   int64   `json:"connect_num"`
	AuthType     int64   `json:"auth_type"`
	ClientType   string  `json:"client_type"`
	Apname       string  `json:"apname"`
	ACGid        int64   `json:"ac_gid"`
	StaticStatus int64   `json:"static_status"`
	DeviceIcon   string  `json:"device_icon"`
	VendorIcon   string  `json:"vendor_icon"`
	Interface    string  `json:"interface"`
	Ppptype      string  `json:"ppptype"`
	IPAddr       string  `json:"ip_addr"`
	UplinkAddr   string  `json:"uplink_addr"`
	UplinkDev    string  `json:"uplink_dev"`
	Termname     string  `json:"termname"`
	TodayTotal   int64   `json:"today_total"`
	VLANID       int64   `json:"vlan_id"`
	Ipv6Gnames   string  `json:"ipv6_gnames"`
	ClientTypeid int64   `json:"client_typeid"`
	ClientVendor string  `json:"client_vendor"`
	ClientModel  string  `json:"client_model"`
	DeviceType   string  `json:"device_type"`
	Username     string  `json:"username"`
	Signal       int64   `json:"signal"`
	Upload       int64   `json:"upload"`
	TotalUp      float64 `json:"total_up"`
	TotalDown    float64 `json:"total_down"`
	SSID         string  `json:"ssid"`
	Frequencies  string  `json:"frequencies"`
	Uptime       string  `json:"uptime"`
	Apmac        string  `json:"apmac"`
	Timestamp    int64   `json:"timestamp"`
	DtalkName    string  `json:"dtalk_name"`
	Ipv4Gnames   string  `json:"ipv4_gnames"`
	MACGnames    string  `json:"mac_gnames"`
	Webid        int64   `json:"webid"`
	LinkAddr     string  `json:"link_addr"`
	Comment      string  `json:"comment"`
	Reject       int64   `json:"reject"`
	Hostname     string  `json:"hostname"`
}

func NewMonitorLanIpAction() *action.Action {
	return &action.Action{
		Action:   "show",
		FuncName: "monitor_lanip",
		Param: map[string]interface{}{
			"TYPE":  "data,total",
			"limit": "0,10000",
		},
	}
}

func NewMonitorLanIPv6Action() *action.Action {
	return &action.Action{
		Action:   "show",
		FuncName: "monitor_lanipv6",
		Param: map[string]interface{}{
			"TYPE":  "data,total",
			"limit": "0,10000",
		},
	}
}

type ShowMonitorLanResult struct {
	Status
	Results struct {
		Data  []LanDeviceInfo `json:"data"`
		Total int64           `json:"total"`
	} `json:"results"`
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

type OnlineUserCount struct {
	Count         int `json:"count"`
	Count2G       int `json:"count_2g"`
	Count5G       int `json:"count_5g"`
	CountWired    int `json:"count_wired"`
	CountWireless int `json:"count_wireless"`
}

type SysStat struct {
	Hostname   string          `json:"hostname"`
	Gwid       string          `json:"gwid"`
	LinkStatus int             `json:"link_status"`
	IPAddr     string          `json:"ip_addr"`
	OnlineUser OnlineUserCount `json:"online_user"`
	Uptime     int             `json:"uptime"`
	Cpu        []string        `json:"cpu"`
	Freq       []string        `json:"freq"`
	Cputemp    []int           `json:"cputemp"`
	Verinfo    Verinfo         `json:"verinfo"`
	Memory     MemoryStat      `json:"memory"`
	Stream     StreamStat      `json:"stream"`
}

type ShowSysStatResult struct {
	Status
	Results struct {
		SysStat SysStat `json:"sysstat"`
	} `json:"results"`
}
