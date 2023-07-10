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
	if len(ret.MachineUUID) == 0 || len(ret.MachineIP) == 0 {
		return fmt.Errorf("机器uuid和ip都不能为空")
	}
	return MySQL().Save(&ret).Error
}

func UpdateStatus(UUID string) error {
	var ret RedisExportTarget
	return MySQL().Model(&ret).Where("uuid=?", UUID).Update("status", global.StatusRemove).Error
}

func GetRedisExporter() ([]RedisExportTarget, error) {
	var ret []RedisExportTarget
	err := MySQL().Where("status=?", global.StatusInstall).Find(&ret).Error
	return ret, err
}
