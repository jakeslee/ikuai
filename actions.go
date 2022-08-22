package ikuai

import (
	"github.com/jakeslee/ikuai/action"
)

// EditIPGroup 修改 IP 组
func (i *IKuai) EditIPGroup(ipGroup action.IPGroup) (*action.Result, error) {
	resp := &action.Result{}

	_, err := i.Run(i.session, action.NewIPGroupEditAction(ipGroup), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ShowIPGroup 取 IP 组列表
func (i *IKuai) ShowIPGroup() (*action.ShowIPGroupResult, error) {
	resp := &action.ShowIPGroupResult{}

	_, err := i.Run(i.session, action.NewIPGroupShowAction(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Switch 执行切换状态
func (i *IKuai) Switch(id string, state action.SwitchState, fn func(string, action.SwitchState) *action.Action) (*action.Result, error) {
	resp := &action.Result{}

	_, err := i.Run(i.session, fn(id, state), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IKuai) ShowMonitorLan() (*action.ShowMonitorResult, error) {
	resp := &action.ShowMonitorResult{}

	_, err := i.Run(i.session, action.NewMonitorLanIpAction(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IKuai) ShowSysStat() (*action.ShowSysStatResult, error) {
	resp := &action.ShowSysStatResult{}

	_, err := i.Run(i.session, action.NewShowSysStatAction(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IKuai) ShowMonitorInterface() (*action.ShowMonitorInterfaceResult, error) {
	resp := &action.ShowMonitorInterfaceResult{}

	_, err := i.Run(i.session, action.NewMonitorInterfaceAction(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
