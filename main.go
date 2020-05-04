package main

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/logs"

	"news/baidu"
	"news/configure"
	"news/serverchan"
	"news/weibo"
)

func init() {
	logsConfig := `{"filename":"./logs/debug.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`
	logs.SetLogger(logs.AdapterFile, logsConfig)
	logs.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

}

func main() {
	duration := time.Hour / 2
	timer := time.NewTimer(duration)

	var cfg configure.Config
	if _, err := toml.DecodeFile("./configure/default.toml", &cfg); err != nil {
		logs.Error("%s", err.Error())
		return
	}
	logs.Debug("scKey: %s", cfg.Key)

	sc := serverchan.NewServerChan(cfg.Key)
	//
	titleWeibo := "微博热搜"
	newsWeibo, _ := weibo.CrawlWeiBoNews()
	sc.PushMsg(titleWeibo, newsWeibo)
	//
	titleBidu := "百度热搜"
	newsBaidu, _ := baidu.CrawlBaiduNews()
	sc.PushMsg(titleBidu, newsBaidu)

	for {
		select {
		case <-timer.C:
			newsWeibo, _ := weibo.CrawlWeiBoNews()
			sc.PushMsg(titleWeibo, newsWeibo)
			//
			newsBaidu, _ := baidu.CrawlBaiduNews()
			sc.PushMsg(titleBidu, newsBaidu)
			//
			timer.Reset(duration)
		default:
			continue
		}
	}
}
