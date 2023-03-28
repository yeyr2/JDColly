package api

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"log"
	"net/url"
	"reptile-test-go/api/cmd"
	"reptile-test-go/api/sql"
	"reptile-test-go/config"
	"reptile-test-go/middleware"
	"strings"
	"time"
)

func GetInfoByJDKey(key string, hots *[]*cmd.Hot) {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)

	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `search.jd.com`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  12,
	})
	if err != nil {
		log.Fatal(err)
	}

	c1 := colly.NewCollector(
		colly.Async(true),
	)
	extensions.RandomUserAgent(c1)

	//设置代理
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1080", "socks5://127.0.0.1:1081")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnHTML("li.gl-item", func(e *colly.HTMLElement) {
		hot := new(cmd.Hot)

		err := e.Unmarshal(hot)
		if err != nil {
			fmt.Println("unmarshal error:", err)
			return
		}
		hot.Img = "https:" + hot.Img
		hot.Url = "https:" + hot.Url
		hot.ProductId = e.Attr("data-sku")

		// 保存图片
		tmp := strings.Split(hot.Img, "/")
		fileName := fmt.Sprintf("images/img%s", tmp[len(tmp)-1])

		c1.OnResponse(func(r *colly.Response) {
			err := r.Save(r.Ctx.Get("file"))
			go middleware.WriteLogFile(fileName, "Images", err)
		})

		c1.OnRequest(func(r *colly.Request) {
			tmp := strings.Split(r.URL.String(), "/")
			fileName := fmt.Sprintf("images/img%s", tmp[len(tmp)-1])
			r.Ctx.Put("file", fileName)
		})

		c1.Visit(hot.Img)

		hot.Img = "http://" + config.Host + "/" + fileName
		hot.Key = key
		*hots = append(*hots, hot)
	})

	keys := url.PathEscape(key)
	err = c.Visit("https://search.jd.com/Search?keyword=" + keys + "&enc=utf-8&wq=" + keys)
	if err != nil {
		log.Fatalln("visit ", err)
	}

	c.Wait()
	c1.Wait()

	go sql.AddShopInfo(hots)
}

func GetInfoByJDKeyBySql(key string, hots *[]*cmd.Hot) {
	sql.GetShopInfoByKey(key, hots)
}
