package logic

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"log"
	"net/url"
	"reptile-test-go/cmd"
	"reptile-test-go/config"
	"reptile-test-go/middleware"
	"reptile-test-go/model"
	"strings"
	"time"
)

func GetInfoByJDKey(key string, hots *[]*cmd.Hot) {
	c := colly.NewCollector(
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)

	c1 := c.Clone()
	err := c1.Limit(&colly.LimitRule{
		DomainRegexp: `img12.360buyimg.com`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  15,
	})
	if err != nil {
		log.Fatal(err)
	}

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

		hot.Img = getClearImgUrl(hot.Img)
		// 保存图片
		tmp := strings.Split(hot.Img, "/")
		fileName := fmt.Sprintf("images/img%s", tmp[len(tmp)-1])

		c1.Visit(hot.Img)

		hot.Img = "http://" + config.Host + "/" + fileName
		hot.Key = key
		*hots = append(*hots, hot)
	})

	c1.OnResponse(func(r *colly.Response) {
		file := r.Ctx.Get("file")
		err := r.Save(file)
		go middleware.WriteLogFile(file, "Images", err)
	})

	c1.OnRequest(func(r *colly.Request) {
		tmp := strings.Split(r.URL.String(), "/")
		fileName := fmt.Sprintf("images/img%s", tmp[len(tmp)-1])
		r.Ctx.Put("file", fileName)
	})

	keys := url.PathEscape(key)
	err = c.Visit("https://search.jd.com/Search?keyword=" + keys + "&enc=utf-8&wq=" + keys)
	if err != nil {
		log.Fatalln("visit ", err)
	}

	c.Wait()
	c1.Wait()

	go sql.AddShopInfo(hots, key)
}

func GetInfoByJDKeyBySql(key string, hots *[]*cmd.Hot) {
	sql.GetShopInfoByKey(key, hots)
}

func getClearImgUrl(img string) string {
	b := strings.Split(img, "/")
	var c strings.Builder
	for i, x := range b {
		if i == 3 {
			c.WriteString("n0")
		} else {
			c.WriteString(x)
		}
		if i != len(b)-1 {
			c.WriteByte('/')
		} else {
			c.WriteString(".avif")
		}
	}
	return c.String()
}
