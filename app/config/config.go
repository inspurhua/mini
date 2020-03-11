package config

import (
	"gopkg.in/ini.v1"
	"log"

	"time"
)

var (
	RunMode   string
	JwtSecret string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	DbType string
	DbUrl  string
	Prefix string
)

func init() {
	var cfg *ini.File
	var err error
	var sec *ini.Section

	cfg, err = ini.Load("app.ini")

	if err != nil {
		log.Fatal("app.ini配置文件有错误:%v", err)
	}
	RunMode = cfg.Section("").Key("app_mode").MustString("debug")
	JwtSecret = cfg.Section("").Key("jwt_secret").MustString("!@)*#)!@U#@*!@!)")

	sec, err = cfg.GetSection("server")
	if err != nil {
		log.Fatal("app.ini配置文件找不到[server]:%v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	sec, err = cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "app.ini配置文件找不到[database]:: %v", err)
	}
	DbType = sec.Key("TYPE").MustString("postgres")
	DbUrl = sec.Key("URL").MustString("")
	Prefix = sec.Key("PREFIX").MustString("sys_")

}
