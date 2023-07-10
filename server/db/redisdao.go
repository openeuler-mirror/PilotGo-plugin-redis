package db

import (
	"fmt"
	"time"

	"openeuler.org/PilotGo/redis-plugin/global"
)

type RedisExportTarget struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MachineUUID string `gorm:"unique" json:"uuid"`
	MachineIP   string `json:"ip"`
	Status      string `json:"status"` //install、remove or error
	Error       string `json:"error"`
	UpdatedAt   time.Time
}

func AddRedisExporter(ret RedisExportTarget) error {
	if len(ret.MachineUUID) == 0 || len(ret.MachineIP) == 0 {
		return fmt.Errorf("机器uuid和ip都不能为空")
	}
	sqlStr := "insert into redis_export_target(machine_uuid, machine_ip, status, error,updated_at) values(?,?,?,?,now()) ON DUPLICATE KEY UPDATE machine_ip=?,status=?,error=?,updated_at=now()"
	res := MySQL().Exec(sqlStr, ret.MachineUUID, ret.MachineIP, ret.Status, ret.Error, ret.MachineIP, ret.Status, ret.Error)
	return res.Error
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
