package service

import (
	"gitee.com/openeuler/PilotGo-plugins/sdk/common"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"

	"openeuler.org/PilotGo/redis-plugin/db"
	"openeuler.org/PilotGo/redis-plugin/global"
	"openeuler.org/PilotGo/redis-plugin/plugin"
)

func FormatData(cmdResults []*client.CmdResult) ([]string, []db.RedisExportTarget, error) {
	ret := []db.RedisExportTarget{}
	monitorTargets := []string{}
	for _, result := range cmdResults {
		d := db.RedisExportTarget{
			MachineUUID: result.MachineUUID,
			MachineIP:   result.MachineIP,
			Status:      "ok",
			Error:       "",
		}

		if result.RetCode != 0 {
			d.Status = global.StatusErroe
			d.Error = result.Stderr
		} else {
			// TODO: add or delete redis exporter to prometheus monitor target here
			// default exporter port :9121
			monitorTargets = append(monitorTargets, result.MachineIP+":9121")
		}

		ret = append(ret, d)
	}
	return monitorTargets, ret, nil
}

func Install(param *common.Batch) ([]db.RedisExportTarget, error) {
	cmd := "yum install -y redis_exporter && systemctl start redis_exporter"

	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return nil, err
	}

	monitorTargets, ret, err := FormatData(cmdResults)
	if err != nil {
		return nil, err
	}
	err = plugin.MonitorTargets(monitorTargets)
	if err != nil {
		return nil, err
	}
	for _, tm := range ret {
		if tm.Status == "ok" {
			tm.Status = global.StatusInstall
		}
		err = db.AddRedisExporter(tm)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	return ret, nil
}

func UnInstall(param *common.Batch) ([]db.RedisExportTarget, error) {
	cmd := "systemctl stop redis_exporter && yum autoremove -y redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return nil, err
	}
	monitorTargets, ret, err := FormatData(cmdResults)
	if err != nil {
		return nil, err
	}
	err = plugin.MonitorTargets(monitorTargets)
	if err != nil {
		return nil, err
	}
	for _, tm := range ret {
		if tm.Status == "ok" {
			tm.Status = global.StatusRemove
		}
		err = db.AddRedisExporter(tm)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	return ret, nil
}

func Restart(param *common.Batch) error {
	cmd := "systemctl restart redis_exporter && systemctl status redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return err
	}
	_, _, err = FormatData(cmdResults)
	return nil
}

func Stop(param *common.Batch) error {
	cmd := "systemctl stop redis_exporter && systemctl status redis_exporter"
	cmdResults, err := global.GlobalClient.RunScript(param, cmd)
	if err != nil {
		return err
	}
	_, _, err = FormatData(cmdResults)
	return nil
}
