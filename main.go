package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// TextReqBody 请求
type TextReqestBody struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   string        // 开发者微信号
	FromUserName string        // 发送方帐号（一个OpenID）
	CreateTime   time.Duration // 消息创建时间 （整型）
	MsgType      string        // 消息类型，文本为text
	Content      string        // 文本消息内容
	MsgId        int           // 消息id，64位整型
}

// 解析XML
func parseTextReqestBody(body []byte) *TextReqestBody {
	requestBody := &TextReqestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

// TextResponseBody 响应
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

// 制作XML
func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextReqestBody{}
	textResponseBody.FromUserName = fromUserName
	textResponseBody.ToUserName = toUserName
	textResponseBody.MsgType = "text"
	textResponseBody.Content = content
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

// WXCallBack 微信回调
func WXCallBack(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("data:\n%s\n", string(data))

	textRequestBody := parseTextReqestBody(data)
	if textRequestBody != nil {

		if textRequestBody.FromUserName == "o8j7ms1TZnSGG9NjuVKFawT66vvU" {
			responseTextBody, _ := makeTextResponseBody(textRequestBody.ToUserName, textRequestBody.FromUserName, "Hello, "+"小口天")
			log.Printf("responseTextBody: %s", string(responseTextBody))
			c.String(200, string(responseTextBody))
		}
	}

	c.String(200, string("success"))
}

// Ping 心跳测试
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "pong"})
}

func main() {
	r := gin.Default()
	r.POST("/wx", WXCallBack)
	r.Any("/", Ping)
	r.Run(":80")
}
