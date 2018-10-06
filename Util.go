package SQLinjCrawler

import (
	"net/url"
	"github.com/valyala/fasthttp"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type responseCallback func(response *fasthttp.Response) bool
type storeCallback func(string)
type preProcess func(url url.URL) string

func doRequest(url string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{ReadTimeout: 10 * time.Second, MaxResponseBodySize: 1024 * 1024}
	client.Do(req, resp)
	return resp
	//bodyBytes := resp.Body()
	//return string(bodyBytes)
}

func Detect(rawUrl string, pre preProcess, parseCallback responseCallback, storeCallback storeCallback) {
	var fetchUrl, err = url.Parse(rawUrl)
	if err != nil {
		return
	}
	urlWithPayload := pre(*fetchUrl)
	if len(urlWithPayload) > 0 {
		if parseCallback(doRequest(urlWithPayload)) {
			storeCallback(rawUrl)
		}
	}
}
