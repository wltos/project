package baidu

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"

	"news/utils"
)

// CrawlBaiduNews 采集百度热搜数据
func CrawlBaiduNews() (string, error) {
	req := httplib.Get("http://top.baidu.com/buzz")
	// 头部
	req.Header("Host", "top.baidu.com")
	req.Header("Connection", "keep-alive")
	req.Header("Cache-Control", "max-age=0")
	req.Header("DNT", "1")
	req.Header("Upgrade-Insecure-Requests", "1")
	req.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header("Referer", "http://top.baidu.com/category?c=513&fr=topindex")
	req.Header("Accept-Encoding", "gzip, deflate")
	req.Header("Accept-Language", "zh,en;q=0.9,zh-CN;q=0.8")
	// 参数
	req.Param("b", "341")
	req.Param("c", "513")
	req.Param("fr", "topcategory_c513")
	// 请求
	resp, err := req.String()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	// 中文编码
	resp = utils.ConvertToString(resp, "gbk", "utf-8")

	// 摘果子
	var desp string
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	domList := dom.Find("table.list-table tr")
	for i := 0; i < domList.Length(); i++ {
		// 表头
		if i == 0 {
			continue
			domTh := domList.Eq(i).Find("th")
			for j := 0; j < domTh.Length(); j++ {
				log.Printf("%s", domTh.Eq(j).Text())
			}
		}

		// 新闻列表
		domTd := domList.Eq(i).Find("td")

		// 排名
		rank := domTd.Eq(0).Find("span").Text()
		log.Println(rank)
		// 关键词
		keyword := domTd.Eq(1).Find("a.list-title").Text()
		kwHref, _ := domTd.Eq(1).Find("a.list-title").Attr("href")
		log.Println(keyword)
		log.Println(kwHref)
		// 新闻
		news := domTd.Eq(2).Find("a").Eq(0).Text()
		newsHref, _ := domTd.Eq(2).Find("a").Eq(0).Attr("href")
		log.Println(news)
		log.Println(newsHref)
		// 视频
		video := domTd.Eq(2).Find("a").Eq(1).Text()
		videoHref, _ := domTd.Eq(2).Find("a").Eq(1).Attr("href")
		log.Println(video)
		log.Println(videoHref)
		// 图片
		image := domTd.Eq(2).Find("a").Eq(2).Text()
		imageHref, _ := domTd.Eq(2).Find("a").Eq(2).Attr("href")
		log.Println(image)
		log.Println(imageHref)
		// 搜索指数
		star := domTd.Eq(3).Find("span").Text()
		log.Println(star)

		desp = desp + fmt.Sprintf("%02s | [%s](%s) | [%s](%s) [%s](%s) [%s](%s) | %s | \n", rank, keyword, kwHref, news, newsHref, video, videoHref, image, imageHref, star)
	}

	head1 := "排名 | 关键词 | 相关链接 | 搜索指数\n"
	head2 := "- | :-: | :-: | -: \n"
	return head1 + head2 + desp, nil
}
