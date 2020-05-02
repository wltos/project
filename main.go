package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"
)

// SCKEY 凭证
const SCKEY string = "fill_your_sckey"

// ServerJiangRsp 消息推送响应
type ServerJiangRsp struct {
	ErrNo   int64  `json:"errno"`
	ErrMsg  string `json:"errmsg"`
	DataSet string `json:"dataset"`
}

// GetWeiBoNews 采集微博热搜数据
func GetWeiBoNews() (string, error) {
	// 打开微博热搜
	req := httplib.Get("https://s.weibo.com/top/summary")
	req.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header("cate", "realtimehot")
	html, err := req.String()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	// 分析数据
	var desp string
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	docTagTr := doc.Find("tbody tr")
	for i := 0; i < docTagTr.Length(); i++ {
		// 排行榜
		index := docTagTr.Eq(i).Find("td").Eq(0).Text()
		if index == "" {
			continue
		}

		// 消息名
		news := docTagTr.Eq(i).Find("a").Text()

		// 链接
		link, _ := docTagTr.Eq(i).Find("a").Attr("href")
		link = fmt.Sprintf("https://s.weibo.com%s", link)

		// 关注数
		star := docTagTr.Eq(i).Find("span").Text()

		msg := fmt.Sprintf("- %2s [%s](%s) %s\n", index, news, link, star)
		desp = desp + msg
	}

	return desp, nil
}

// SendMsg 微信推送
func SendMsg(text, desp string) error {
	rawURL := fmt.Sprintf("http://sc.ftqq.com/%s.send", SCKEY)

	serverJiang := httplib.Post(rawURL)
	serverJiang.Param("text", text)
	serverJiang.Param("desp", desp)
	resp, err := serverJiang.String()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	var tmp ServerJiangRsp
	if err := json.Unmarshal([]byte(resp), &tmp); err != nil {
		log.Println(err.Error())
		return err
	}

	if tmp.ErrNo == 0 || tmp.ErrMsg == "success" {
		return nil
	}
	return errors.New("发送失败")
}

func main() {
	duration := time.Hour
	timer := time.NewTimer(duration)

	msg, _ := GetWeiBoNews()
	SendMsg("微博热搜", msg)

	for {
		select {
		case <-timer.C:
			// 准备数据
			msg, _ := GetWeiBoNews()
			// 微信推送
			if err := SendMsg("微博热搜", msg); err != nil {
				log.Println(err.Error())
			}
			// 复位
			timer.Reset(duration)
		default:
			continue
		}
	}
}
