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

func GetCommentById(id string, comment *JDComment) bool {
	chans := make(chan *[]Comments, 30)

	for i := 0; i < 20; i++ {
		go AddComment(id, i, chans)

		// 防止爬得过快
		time.Sleep(100 * time.Millisecond)
	}

	for i := 0; i < 20; i++ {
		(*comment).Comments = append((*comment).Comments, *(<-chans)...)
	}

	return true
}

func AddComment(id string, i int, chans chan *[]Comments) {

	newUrl := GetCommentUrl(id, strconv.Itoa(i))
	result := SendHttp(newUrl)

	var tmp JDComment
	if len(*result) == 0 {
		return
	}
	err := json.Unmarshal(*result, &tmp)
	if err != nil {
		log.Fatalln(err)
	}

	chans <- &tmp.Comments
}

func GetCommentUrl(productId string, page string) string {
	return "https://club.jd.com/comment/productPageComments.action?productId=" + productId + "&score=0&sortType=5&page=" + page + "&pageSize=10&isShadowSku=0&fold=1"
}

func SendHttp(urls string) (body *[]byte) {
	body = new([]byte)
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	extensions.RandomUserAgent(c)

	rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:1080", "http://127.0.0.1:1081")
	if err != nil {
		log.Fatal(err)
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
	})

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		body = &r.Body
	})

	c.Visit(urls)

	c.Wait()

	return body
}

func AnalyzeGetComments(comment *JDComment, analyze *AnalyzeComment) {

}

func WordCloudAnalysis(comment *JDComment, analyze *AnalyzeComment) {

}
