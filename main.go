package main

import (
	"time"

	"github.com/astaxie/beego/logs"

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

	// 初次提交
	text := "微博热搜"
	desp, _ := weibo.CrawlWeiBoNews()
	sc := serverchan.NewServerChan("FILL_YOUR_SCKEY")
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
