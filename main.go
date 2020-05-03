package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/astaxie/beego/httplib"
)

func main() {
	// 拉取数据
	data, err := lsjz(1)
	if err != nil {
		log.Fatal(err)
	}

	// 转换数据
	var buf rsp
	if err := json.Unmarshal([]byte(data), &buf); err != nil {
		log.Fatal(err)
	}
	log.Printf("总条数: %d", len(buf.Data.LSJZList))

	f := excelize.NewFile()
	// 写表头
	index := f.NewSheet("历史净值明细")
	f.SetCellValue("历史净值明细", "A1", "净值日期")
	f.SetCellValue("历史净值明细", "B1", "单位净值")
	f.SetCellValue("历史净值明细", "C1", "累计净值")
	f.SetCellValue("历史净值明细", "D1", "日增长率")
	f.SetCellValue("历史净值明细", "E1", "申购状态")
	f.SetCellValue("历史净值明细", "F1", "赎回状态")
	f.SetCellValue("历史净值明细", "G1", "分红送配")

	// 写数据
	for k, v := range buf.Data.LSJZList {
		i := fmt.Sprintf("%d", k+2)
		f.SetCellValue("历史净值明细", "A"+i, v.Fsrq)
		f.SetCellValue("历史净值明细", "B"+i, v.Dwjz)
		f.SetCellValue("历史净值明细", "C"+i, "")
		f.SetCellValue("历史净值明细", "D"+i, v.Jzzzl+"%")
		f.SetCellValue("历史净值明细", "E"+i, v.Sgzt)
		f.SetCellValue("历史净值明细", "F"+i, v.Shzt)
		f.SetCellValue("历史净值明细", "G"+i, "")
	}

	// 保存
	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(index)
	if err := f.SaveAs("易方达银行分级(161121)_基金历史净值_基金档案.xlsx"); err != nil {
		println(err.Error())
	}
}

// 接住数据
type rsp struct {
	Data struct {
		Feature  string `json:"Feature"`
		FundType string `json:"FundType"`
		LSJZList []struct {
			Actualsyi string      `json:"ACTUALSYI"`
			Dtype     interface{} `json:"DTYPE"`
			Dwjz      string      `json:"DWJZ"` // 单位净值
			Fhfcbz    string      `json:"FHFCBZ"`
			Fhfcz     string      `json:"FHFCZ"`
			Fhsp      string      `json:"FHSP"`
			Fsrq      string      `json:"FSRQ"` // 净值日期
			Jzzzl     string      `json:"JZZZL"`
			Ljjz      string      `json:"LJJZ"`
			Navtype   string      `json:"NAVTYPE"`
			Sdate     interface{} `json:"SDATE"`
			Sgzt      string      `json:"SGZT"` // 申购状态
			Shzt      string      `json:"SHZT"` // 赎回状态
		} `json:"LSJZList"`
		SYType    interface{} `json:"SYType"`
		IsNewType bool        `json:"isNewType"`
	} `json:"Data"`
	ErrCode    int64       `json:"ErrCode"`
	ErrMsg     interface{} `json:"ErrMsg"`
	Expansion  interface{} `json:"Expansion"`
	PageIndex  int64       `json:"PageIndex"`
	PageSize   int64       `json:"PageSize"`
	TotalCount int64       `json:"TotalCount"`
}

// 模拟用户请求
func lsjz(pageIndex int64) (string, error) {
	req := httplib.Get("http://api.fund.eastmoney.com/f10/lsjz")
	req.Header("Host", "api.fund.eastmoney.com")
	req.Header("Connection", "keep-alive")
	req.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header("DNT", "1")
	req.Header("Accept", "*/*")
	req.Header("Referer", "http://fundf10.eastmoney.com/jjjz_161121.html")
	req.Header("Accept-Encoding", "gzip, deflate")
	req.Header("Accept-Language", "zh,en;q=0.9,zh-CN;q=0.8")
	req.Param("callback", "jQuery18305110095191834001_1588514143393")
	req.Param("fundCode", "161121")
	req.Param("pageIndex", fmt.Sprintf("%d", pageIndex))
	req.Param("pageSize", "1203") // 全量请求
	req.Param("startDate", "")
	req.Param("endDate", "")
	req.Param("_", fmt.Sprintf("%d", time.Now().UnixNano()/1e6))

	// 调试
	// req.SetProxy(func(req *http.Request) (*url.URL, error) {
	// 	u, _ := url.ParseRequestURI("http://127.0.0.1:8888")
	// 	return u, nil
	// })

	resp, err := req.String()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	resp = resp[strings.Index(resp, "(")+1 : strings.Index(resp, ")")]
	log.Println(resp)
	return resp, nil
}
