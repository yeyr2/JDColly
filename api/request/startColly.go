package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"log"
	"net/http"
	"net/url"
	"reptile-test-go/api/cmd"
	"reptile-test-go/setting"
	"strings"
	"time"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

func StartColly(con *gin.Context) {
	key := con.Query("key")

	var hots []*cmd.Hot

	getInfoByJDKey(key, &hots)

	origin := con.GetHeader("Origin")

	if origin != "" {
		con.Header("Access-Control-Allow-Origin", origin)
	}
	con.Header("Access-Control-Allow-Methods", "*")

	headers := con.GetHeader("Access-Control-Allow-Headers")
	log.Println(headers)
	if headers != "" {
		con.Header("Access-Control-Allow-Headers", headers)
	}

	con.Header("Access-Control-Max-Age", "3600")

	con.Header("Access-Control-Allow-Credentials", "true")

	con.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      hots,
	})
}

func getInfoByJDKey(key string, hots *[]*cmd.Hot) {
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
			if err != nil {
				log.Printf("saving %s failed:%v\n", fileName, err)
			} else {
				log.Printf("saving %s success\n", fileName)
			}
		})

		c1.OnRequest(func(r *colly.Request) {
			tmp := strings.Split(r.URL.String(), "/")
			fileName := fmt.Sprintf("images/img%s", tmp[len(tmp)-1])
			r.Ctx.Put("file", fileName)
		})

		c1.Visit(hot.Img)

		hot.Img = "http://" + setting.Host + ":" + setting.Port + "/" + fileName
		*hots = append(*hots, hot)
	})

	key = url.PathEscape(key)
	err = c.Visit("https://search.jd.com/Search?keyword=" + key + "&enc=utf-8&wq=" + key)
	if err != nil {
		log.Fatalln("visit ", err)
	}

	c.Wait()
	c1.Wait()
}
