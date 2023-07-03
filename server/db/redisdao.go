package db

import (
	"fmt"
	"time"

	"openeuler.org/PilotGo/redis-plugin/global"
)

type RedisExportTarget struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MachineUUID string `json:"uuid"`
	MachineIP   string `json:"ip"`
	Status      string `json:"status"` //install、remove or error
	Error       string `json:"error"`
	UpdatedAt   time.Time
}

func AddRedisExporter(ret RedisExportTarget) error {
	if len(ret.MachineUUID) == 0 {
		return fmt.Errorf("机器不能为空")
	}
	return global.GlobalDB.Save(&ret).Error
}

func GetRedisExporter() ([]RedisExportTarget, error) {
	var ret []RedisExportTarget
	err := global.GlobalDB.Where("status=?", global.StatusInstall).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
