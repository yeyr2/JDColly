package api

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"log"
	"strconv"
	"time"
)

var proxys = []string{"socks5://127.0.0.1:1080", "socks5://127.0.0.1:1081"}

func GetCommentById(id string, lastTime int64, comment *JDComment) bool {

	GetTotalPages(id, comment, lastTime)
	time.Sleep(1 * time.Second)

	pages := comment.MaxPage

	if pages <= 0 {
		return false
	} else if pages == 1 { // 已经获取过
		return true
	} else if pages > 30 { // 临时的
		pages = 30
	}

	err := SendHttp(id, comment, pages, lastTime)
	if err != nil {
		log.Println("err:", err)
		return false
	}

	return true
}

func SendHttp(productId string, comment *JDComment, pages int, lastTime int64) (err error) {
	chans := make(chan *[]Comments, pages-1)

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)

	//rp, err := proxy.RoundRobinProxySwitcher("http://117.74.65.215:8082", "http://103.151.60.204:80", "http://188.165.227.155:5397")
	rp, err := proxy.RoundRobinProxySwitcher(proxys...)
	if err != nil {
		log.Fatal("proxy:", err)
	}
	c.SetProxyFunc(rp)
	//c.SetProxyFunc(http.ProxyFromEnvironment)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Content-Type", "application/x-www-form-urlencoded;charset=GBK")
		r.Headers.Add("cookie", " __jdu=16736026655301847779361; unpl=JF8EAKxnNSttURhTUhIBGhQYTFpWW1sMTkdRPGQAUloKTAQFGgUdERd7XlVdXxRKFB9sYhRVVFNPXA4bAisSEXtfVl5dAUgTAW9XNVNdUQYVV1YyGBIgS1xkXloOSRACbGACUFlaSlQMHgQdFBhJWlRuXDhMFwpfVzVRXVlKXQMbBxkWIEptVl9cDUMfBWxuB2QWNkoZBRwEGRURSFpTWlkKShcKamEDUlVaTFQ1GjIY; __jdv=76161171|www.google.com.hk|-|referral|-|1678705291683; areaId=27; PCSYCityID=CN_610000_610100_0; shshshfpa=3c88ad85-a0f8-ff4e-abb1-6a3e3dd753c9-1678705294; shshshfpx=3c88ad85-a0f8-ff4e-abb1-6a3e3dd753c9-1678705294; shshshfpb=hmCm-O48oKw55cFHxuf-HAg; ipLoc-djd=27-2376-50232-53749; jwotest_product=99; __jdc=122270672; 3AB9D23F7A4B3CSS=jdd03QRZY3PGTFEQEW7PUUUZ2O52IZZ3YVV7VNSIQVAEVK4FFJ4KBCB45VKF74BWGMWKJRV3FGSQHFETTWFI4XX4FESQPLUAAAAMG5GYJGKAAAAAAC3NQNNFBZP4ANYX; shshshfp=99dcccf4d29f2170b5fc295bc36dce06; JSESSIONID=F4A0F2918C017C70E717AF256385D4A3.s1; token=3e9b9793564885eedab9a1e719849140,2,932759; __tk=4UGFkuWE4UowAYbtAVeykuXpAVPukrPqkUAq4ugE4r4z4Vj0Bc4xAw,2,932759; jsavif=1; shshshsID=4ba1dcc45916dd6359b6a286d8e6f364_1_1678966889076; __jda=122270672.16736026655301847779361.1673602666.1678963121.1678966889.17; __jdb=122270672.1.16736026655301847779361|17.1678966889; 3AB9D23F7A4B3C9B=QRZY3PGTFEQEW7PUUUZ2O52IZZ3YVV7VNSIQVAEVK4FFJ4KBCB45VKF74BWGMWKJRV3FGSQHFETTWFI4XX4FESQPLU")
		r.Headers.Add("authority", " club.jd.com")
		r.Headers.Add("accept", " */*")
		r.Headers.Add("accept-language", "zh-CN,zh;q=0.9")
		r.Headers.Add("referer", "https://item.jd.com/")
		log.Println(r.URL)
	})

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		go JsonBody(&r.Body, lastTime, chans)
	})

	for i := 1; i < pages; i++ {
		urls := GetCommentUrl(productId, strconv.Itoa(i))
		err = c.Visit(urls)
		if err != nil {
			return err
		}

		// 防止爬得过快
		time.Sleep(1 * time.Second)
	}

	for i := 1; i < pages; i++ {
		(*comment).Comments = append((*comment).Comments, *(<-chans)...)
	}

	c.Wait()

	return nil
}

func GetCommentUrl(productId string, page string) string {
	return "https://club.jd.com/comment/productPageComments.action?productId=" + productId + "&score=0&sortType=5&page=" + page + "&pageSize=10&isShadowSku=0&fold=1"
}

func JsonBody(body *[]byte, lastTime int64, chans chan *[]Comments) {
	var tmp JDComment
	if len(*body) == 0 {
		return
	}

	err := json.Unmarshal(*body, &tmp)
	if err != nil {
		log.Println(err)
		return
	}

	DeleteCommentByLastTime(&tmp.Comments, lastTime)

	chans <- &tmp.Comments
}

func GetTotalPages(id string, comment *JDComment, lastTime int64) {
	urls := GetCommentUrl(id, "0")

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	extensions.RandomUserAgent(c)

	rp, err := proxy.RoundRobinProxySwitcher(proxys...)
	if err != nil {
		log.Fatal("proxy:", err)
	}
	c.SetProxyFunc(rp)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Content-Type", "application/x-www-form-urlencoded;charset=GBK")
		r.Headers.Add("cookie", " __jdu=16736026655301847779361; unpl=JF8EAKxnNSttURhTUhIBGhQYTFpWW1sMTkdRPGQAUloKTAQFGgUdERd7XlVdXxRKFB9sYhRVVFNPXA4bAisSEXtfVl5dAUgTAW9XNVNdUQYVV1YyGBIgS1xkXloOSRACbGACUFlaSlQMHgQdFBhJWlRuXDhMFwpfVzVRXVlKXQMbBxkWIEptVl9cDUMfBWxuB2QWNkoZBRwEGRURSFpTWlkKShcKamEDUlVaTFQ1GjIY; __jdv=76161171|www.google.com.hk|-|referral|-|1678705291683; areaId=27; PCSYCityID=CN_610000_610100_0; shshshfpa=3c88ad85-a0f8-ff4e-abb1-6a3e3dd753c9-1678705294; shshshfpx=3c88ad85-a0f8-ff4e-abb1-6a3e3dd753c9-1678705294; shshshfpb=hmCm-O48oKw55cFHxuf-HAg; ipLoc-djd=27-2376-50232-53749; jwotest_product=99; __jdc=122270672; 3AB9D23F7A4B3CSS=jdd03QRZY3PGTFEQEW7PUUUZ2O52IZZ3YVV7VNSIQVAEVK4FFJ4KBCB45VKF74BWGMWKJRV3FGSQHFETTWFI4XX4FESQPLUAAAAMG5GYJGKAAAAAAC3NQNNFBZP4ANYX; shshshfp=99dcccf4d29f2170b5fc295bc36dce06; JSESSIONID=F4A0F2918C017C70E717AF256385D4A3.s1; token=3e9b9793564885eedab9a1e719849140,2,932759; __tk=4UGFkuWE4UowAYbtAVeykuXpAVPukrPqkUAq4ugE4r4z4Vj0Bc4xAw,2,932759; jsavif=1; shshshsID=4ba1dcc45916dd6359b6a286d8e6f364_1_1678966889076; __jda=122270672.16736026655301847779361.1673602666.1678963121.1678966889.17; __jdb=122270672.1.16736026655301847779361|17.1678966889; 3AB9D23F7A4B3C9B=QRZY3PGTFEQEW7PUUUZ2O52IZZ3YVV7VNSIQVAEVK4FFJ4KBCB45VKF74BWGMWKJRV3FGSQHFETTWFI4XX4FESQPLU")
		r.Headers.Add("authority", " club.jd.com")
		r.Headers.Add("accept", " */*")
		r.Headers.Add("accept-language", "zh-CN,zh;q=0.9")
		r.Headers.Add("referer", "https://item.jd.com/")
		log.Println(r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if len(r.Body) == 0 {
			return
		}

		err = json.Unmarshal(r.Body, comment)
		if err != nil {
			log.Println(err)
		}

		DeleteCommentByLastTime(&comment.Comments, lastTime)
	})
	c.Visit(urls)

	c.Wait()
}

func DeleteCommentByLastTime(comments *[]Comments, lastTime int64) {
	length := len(*comments)
	for i := 0; i < length; i++ {
		t, err := time.Parse("2006-01-02 15:04:05", (*comments)[i].ReferenceTime)
		if err != nil {
			log.Println(err)
			continue
		}
		if t.Unix() < lastTime {
			*comments = append((*comments)[:i], (*comments)[i+1:]...)
			i--
			length--
		}
	}
}
