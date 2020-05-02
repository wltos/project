package weibo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/wltos/httplib"
)

// CrawlWeiBoNews 采集微博热搜数据
func CrawlWeiBoNews() (string, error) {
	req := httplib.Get("https://s.weibo.com/top/summary")
	req.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header("cate", "realtimehot")
	html, err := req.String()
	if err != nil {
		logs.Error("%s", err.Error())
		return "", err
	}

	var desp string
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	docTagTr := doc.Find("tbody tr")

	for i := 0; i < docTagTr.Length(); i++ {
		index := docTagTr.Eq(i).Find("td").Eq(0).Text()
		if index == "" {
			continue
		}
		news := docTagTr.Eq(i).Find("a").Text()
		link, _ := docTagTr.Eq(i).Find("a").Attr("href")
		link = fmt.Sprintf("https://s.weibo.com%s", link)
		star := docTagTr.Eq(i).Find("span").Text()

		desp = desp + fmt.Sprintf("- %2s [%s](%s) %s\n", index, news, link, star)
	}

	return desp, nil
}
