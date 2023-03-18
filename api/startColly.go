package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
	"net/http"
	"net/url"
	"time"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

func StartColly(con *gin.Context) {
	key := con.Query("key")
	var hots []*Hot

	getInfoByJDKey(key, &hots)

	con.JSON(http.StatusOK, hots)
}

func getInfoByJDKey(key string, hots *[]*Hot) {
	c := colly.NewCollector(
		//colly.UserAgent(UserAgent),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)

	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `search.jd.com`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  24,
	})
	if err != nil {
		log.Fatal(err)
	}

	c1 := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	//设置代理
	//rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.SetProxyFunc(rp)

	c.OnHTML("li.gl-item", func(e *colly.HTMLElement) {
		hot := new(Hot)

		err := e.Unmarshal(hot)
		if err != nil {
			fmt.Println("unmarshal error:", err)
			return
		}
		hot.Img = "https:" + hot.Img
		hot.Url = "https:" + hot.Url
		hot.ProductId = e.Attr("data-sku")

		hot.Status = GetCommentByUrl(hot)

		fileName := fmt.Sprintf("images/img%s.jpg", hot.ProductId)
		c1.OnResponse(func(r *colly.Response) {
			err := r.Save(fileName)
			if err != nil {
				log.Printf("saving %s failed:%v\n", fileName, err)
			} else {
				log.Printf("saving %s success\n", fileName)
			}
		})

		hot.Img = "http://127.0.0.1:9090/" + fileName
		c1.Visit(hot.Img)

		*hots = append(*hots, hot)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response:", len(r.Body))
	})

	key = url.PathEscape(key)
	err = c.Visit("https://search.jd.com/Search?keyword=" + key + "&enc=utf-8&wq=" + key)
	if err != nil {
		log.Fatalln("visit ", err)
	}

	c.Wait()
	c1.Wait()
}
