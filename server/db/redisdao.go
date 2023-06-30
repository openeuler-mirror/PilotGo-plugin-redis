package db

import (
	"fmt"
	"time"

	"openeuler.org/PilotGo/redis-plugin/global"
)

type RedisExportTarget struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID      string `json:"uuid"`
	Status    string `json:"status"` //install or remove
	UpdatedAt time.Time
}

func AddRedisExporter(ret RedisExportTarget) error {
	if len(ret.UUID) == 0 {
		return fmt.Errorf("机器不能为空")
	}
	return global.GlobalDB.Save(&ret).Error
}

func GetRedisExporter() ([]RedisExportTarget, error) {
	var ret []RedisExportTarget
	err := global.GlobalDB.Where("status=?", "install").Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
