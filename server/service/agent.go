package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"

	"openeuler.org/PilotGo/redis-plugin/db"
	"openeuler.org/PilotGo/redis-plugin/global"
)

func FormatData(cmdResults []*client.CmdResult) ([]db.RedisExportTarget, error) {
	ret := []db.RedisExportTarget{}
	for _, result := range cmdResults {
		d := db.RedisExportTarget{
			MachineUUID: result.MachineUUID,
			MachineIP:   result.MachineIP,
			Status:      "ok",
			Error:       "",
		}
		if result.RetCode != 0 {
			d.Status = global.StatusError
			d.Error = result.Stderr
		}
		ret = append(ret, d)
	}
	return ret, nil
}

func Install(param *common.Batch) ([]db.RedisExportTarget, []db.RedisExportTarget, error) {
	cmd := "yum install -y redis_exporter && systemctl start redis_exporter"
	cmdResults, err := global.GlobalClient.RunCommand(param, cmd)
	if err != nil {
		return nil, nil, err
	}
	ret, err := FormatData(cmdResults)
	if err != nil {
		return nil, nil, err
	}
	var failRet []db.RedisExportTarget
	for _, tm := range ret {
		if tm.Status == "ok" {
			tm.Status = global.StatusInstall
		}
		err = db.AddRedisExporter(tm)
		if err != nil {
			failRet = append(failRet, tm)
			logger.Error(err.Error())
		}
	}
	return ret, failRet, nil
}

func UnInstall(param *common.Batch) ([]db.RedisExportTarget, []db.RedisExportTarget, error) {
	cmd := "systemctl stop redis_exporter && yum autoremove -y redis_exporter"
	cmdResults, err := global.GlobalClient.RunCommand(param, cmd)
	if err != nil {
		return nil, nil, err
	}
	ret, err := FormatData(cmdResults)
	if err != nil {
		return nil, nil, err
	}
	var failRet []db.RedisExportTarget
	for _, tm := range ret {
		if tm.Status == "ok" {
			err = db.UpdateStatus(tm.MachineUUID)
			if err != nil {
				failRet = append(failRet, tm)
				logger.Error(err.Error())
			}
		}
	}
	return ret, failRet, nil
}

func Restart(param *common.Batch) ([]db.RedisExportTarget, error) {
	_, err := Stop(param)
	if err != nil {
		return nil, err
	}
	cmdResults, err := global.GlobalClient.StartService(param, "redis_exporter")
	if err != nil {
		return nil, err
	}

	ret := []db.RedisExportTarget{}
	for _, result := range cmdResults {
		d := db.RedisExportTarget{
			MachineUUID: result.MachineUUID,
			MachineIP:   result.MachineIP,
			Status:      result.ServiceSample.ServiceActiveStatus,
			Error:       "",
		}
		ret = append(ret, d)
	}

	return ret, err
}

func Stop(param *common.Batch) ([]db.RedisExportTarget, error) {
	cmdResults, err := global.GlobalClient.StopService(param, "redis_exporter")
	if err != nil {
		return nil, err
	}

	ret := []db.RedisExportTarget{}
	for _, result := range cmdResults {
		d := db.RedisExportTarget{
			MachineUUID: result.MachineUUID,
			MachineIP:   result.MachineIP,
			Status:      result.ServiceSample.ServiceActiveStatus,
			Error:       "",
		}
		ret = append(ret, d)
	}

	return ret, err
}

func ServiceStatus(param *common.Batch) ([]db.RedisExportTarget, error) {
	cmdResults, err := global.GlobalClient.ServiceStatus(param, "redis_exporter")
	if err != nil {
		return nil, err
	}

	ret := []db.RedisExportTarget{}
	for _, result := range cmdResults {
		d := db.RedisExportTarget{
			MachineUUID: result.MachineUUID,
			MachineIP:   result.MachineIP,
			Status:      result.ServiceSample.ServiceActiveStatus,
			Error:       "",
		}
		ret = append(ret, d)
	}

	return ret, err
}

func GetRedisExporterIp() ([]string, error) {
	ret, err := db.GetRedisExporter()
	if err != nil {
		return nil, err
	}
	var targets []string
	for _, tm := range ret {
		target := tm.MachineIP + ":9121"
		targets = append(targets, target)
	}
	return targets, nil
}
