package funcs

import (
	"github.com/cihub/seelog"
	"github.com/yongbostudio/satagent/satagent/src/common"
)

func StartCloudMonitor() {
	seelog.Info("[func:StartCloudMonitor] ", "starting run StartCloudMonitor ")
	_, err := common.SaveCloudConfig(common.Cfg.Mode["Endpoint"])
	if err != nil {
		seelog.Error("[func:StartCloudMonitor] Cloud Monitor Error", err)
		return
	}
	saveerr := common.SaveConfig()
	if saveerr != nil {
		seelog.Error("[func:StartCloudMonitor] Save Cloud Config Error", err)
		return
	}
	seelog.Info("[func:StartCloudMonitor] ", "StartCloudMonitor finish ")

}
