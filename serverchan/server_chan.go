package serverchan

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

// ServerChan Server酱
type ServerChan struct {
	scKey string
}

// NewServerChan 创建对象
func NewServerChan(sckey string) *ServerChan {
	return &ServerChan{scKey: sckey}
}

// PushMsg 消息推送
func (sc *ServerChan) PushMsg(text, desp string) error {
	rawURL := fmt.Sprintf("http://sc.ftqq.com/%s.send", sc.scKey)
	req := httplib.Post(rawURL)
	req.Param("text", text)
	req.Param("desp", desp)
	resp, err := req.String()
	if err != nil {
		logs.Error("%s", err.Error())
		return err
	}

	var result rsp
	if err := json.Unmarshal([]byte(resp), &result); err != nil {
		logs.Error("%s", err.Error())
		return err
	}

	if result.errNo == 0 || result.errMsg == "success" {
		logs.Debug("标题: %s", text)
		logs.Debug("正文: %s", desp)
		return nil
	}

	errMsg := fmt.Sprintf("推送失败: %s", result.errMsg)
	logs.Debug(errMsg)
	return errors.New(errMsg)
}
