package main

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/logs"

	"weibotop/configure"
	"weibotop/serverchan"
	"weibotop/weibo"
)

func init() {
	logsConfig := `{"filename":"./logs/debug.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`
	logs.SetLogger(logs.AdapterFile, logsConfig)
	logs.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

}

func main() {
	duration := time.Hour
	timer := time.NewTimer(duration)

	var cfg configure.Config
	if _, err := toml.DecodeFile("./configure/default.toml", &cfg); err != nil {
		logs.Error("%s", err.Error())
		return
	}
	logs.Debug("scKey: %s", cfg.Key)

	// 初次提交
	text := "微博热搜"
	desp, _ := weibo.CrawlWeiBoNews()
	sc := serverchan.NewServerChan(cfg.Key)
	sc.PushMsg(text, desp)

	for {
		select {
		case <-timer.C:
			desp, _ := weibo.CrawlWeiBoNews()
			if err := sc.PushMsg(text, desp); err != nil {
				logs.Error("%s", err.Error())
			}

			timer.Reset(duration)
		default:
			continue
		}
	}
}
