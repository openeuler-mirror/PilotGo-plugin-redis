package db

import (
	"database/sql"
	"fmt"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"openeuler.org/PilotGo/redis-plugin/config"
	"openeuler.org/PilotGo/redis-plugin/global"
)

var Url string

type MysqlManager struct {
	ip       string
	port     int
	username string
	password string
	dbname   string
	db       *gorm.DB
}

func ensureDatabase(conf *config.MysqlDBInfo) {
	Url = fmt.Sprintf("%s:%s@(%s:%d)/?charset=utf8mb4&parseTime=true",
		conf.UserName,
		conf.Password,
		conf.HostName,
		conf.Port)
	db, err := gorm.Open(mysql.Open(Url))
	if err != nil {
		logger.Error(err.Error())
	}

	creatDataBase := "CREATE DATABASE IF NOT EXISTS " + conf.DataBase + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci"
	db.Exec(creatDataBase)
}

func MysqldbInit(conf *config.MysqlDBInfo) error {
	// 检查数据库是否存在，不存在则创建
	ensureDatabase(conf)

	m := &MysqlManager{
		ip:       conf.HostName,
		port:     conf.Port,
		username: conf.UserName,
		password: conf.Password,
		dbname:   conf.DataBase,
	}
	Url = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		m.username,
		m.password,
		m.ip,
		m.port,
		m.dbname)
	var err error
	m.db, err = gorm.Open(mysql.Open(Url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	global.GlobalDB = m.db

	var db *sql.DB
	if db, err = m.db.DB(); err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	global.GlobalDB.AutoMigrate(&RedisExportTarget{})
	return nil
}
