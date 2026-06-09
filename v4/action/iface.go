package action

import "github.com/jakeslee/ikuai/action"

type MonitorIfaceType string

const (
	IfaceTypeInterfaceStream MonitorIfaceType = "iface_stream"
	IfaceTypeEtherInfo       MonitorIfaceType = "ether_info"
	IfaceTypeSnapshoot       MonitorIfaceType = "snapshoot"
	IfaceTypeInterfaceCheck  MonitorIfaceType = "iface_check"
)

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

type SnapshootLAN struct {
	ID        int64    `json:"id"`
	Comment   string   `json:"comment"`
	Interface string   `json:"interface"`
	Bandmode  int64    `json:"bandmode"`
	Linkmode  int64    `json:"linkmode"`
	MAC       string   `json:"mac"`
	Member    []string `json:"member"`
	IPAddr    string   `json:"ip_addr"`
	Netmask   string   `json:"netmask"`
}

type SnapshootWAN struct {
	ID             int64    `json:"id"`
	Comment        string   `json:"comment"`
	Interface      string   `json:"interface"`
	MAC            string   `json:"mac"`
	Member         []string `json:"member"`
	Bandmode       int64    `json:"bandmode"`
	DefaultRoute   int64    `json:"default_route"`
	Internet       int64    `json:"internet"`
	IPAddr         string   `json:"ip_addr"`
	Netmask        string   `json:"netmask"`
	Gateway        string   `json:"gateway"`
	Dns1           string   `json:"dns1"`
	Dns2           string   `json:"dns2"`
	CountStatic    int64    `json:"count_static"`
	CountDHCP      int64    `json:"count_dhcp"`
	CountPppoe     int64    `json:"count_pppoe"`
	CountCheckFail int64    `json:"count_check_fail"`
	Updatetime     int64    `json:"updatetime"`
	CheckRes       int64    `json:"check_res"`
	Errmsg         string   `json:"errmsg"`
	Power          string   `json:"power"`
	ISP            string   `json:"isp"`
	Imei           string   `json:"imei"`
	Qnw            string   `json:"qnw"`
	Ccid           string   `json:"ccid"`
	Isnr           string   `json:"isnr"`
	Pcid           string   `json:"pcid"`
}

type Eth struct {
	Driver    string `json:"driver"`
	Type      string `json:"type"`
	MAC       string `json:"mac"`
	Link      int64  `json:"link"`
	Speed     int64  `json:"speed"`
	Duplex    int64  `json:"duplex"`
	Model     string `json:"model"`
	Interface string `json:"interface"`
	Lock      int64  `json:"lock"`
	Bindmod   int64  `json:"bindmod"`
}

type ShowMonitorInterfaceResult struct {
	Status
	Results struct {
		IfaceStream  []IfaceStream  `json:"iface_stream,omitempty"`
		IfaceCheck   []IfaceCheck   `json:"iface_check,omitempty"`
		EtherInfo    map[string]Eth `json:"ether_info,omitempty"`
		SnapshootLAN []SnapshootLAN `json:"snapshoot_lan,omitempty"`
		SnapshootWAN []SnapshootWAN `json:"snapshoot_wan,omitempty"`
	} `json:"results"`
}

func NewMonitorInterfaceAction(typeNames ...MonitorIfaceType) *action.Action {
	return &action.Action{
		Action:   "show",
		FuncName: "monitor_iface",
		Param: map[string]interface{}{
			"TYPE": Join(typeNames, ",", IfaceTypeInterfaceStream),
		},
	}
}
