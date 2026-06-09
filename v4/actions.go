package v4

import (
	"github.com/jakeslee/ikuai/v4/action"
)

func (i *IKuaiV4) ShowMonitorLan() (*action.ShowMonitorLanResult, error) {
	resp := &action.ShowMonitorLanResult{}

	_, err := i.Run(action.NewMonitorLanIpAction(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IKuaiV4) ShowMonitorLanIPv6() (*action.ShowMonitorLanResult, error) {
	resp := &action.ShowMonitorLanResult{}

	_, err := i.Run(action.NewMonitorLanIPv6Action(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IKuaiV4) ShowSysStat() (*action.ShowSysStatResult, error) {
	resp := &action.ShowSysStatResult{}

	_, err := i.Run(action.NewShowHomepageAction(action.HomepageTypeSysStat), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IKuaiV4) ShowMonitorInterface() (*action.ShowMonitorInterfaceResult, error) {
	resp := &action.ShowMonitorInterfaceResult{}

	_, err := i.Run(action.NewMonitorInterfaceAction(action.IfaceTypeInterfaceCheck, action.IfaceTypeInterfaceStream), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
