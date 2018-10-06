package SQLinjCrawler

import (
	"net/url"
	"github.com/valyala/fasthttp"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type storeCallback func(url url.URL, in PluginIn)

func DefaultRequest(url string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{ReadTimeout: 10 * time.Second, MaxResponseBodySize: 1024 * 1024}
	client.Do(req, resp)
	return resp
	//bodyBytes := resp.Body()
	//return string(bodyBytes)
}

func Detect(rawUrl string, plugin PluginIn, storeCallback storeCallback) {
	var fetchUrl, err = url.Parse(rawUrl)
	if err != nil {
		return
	}
	urlWithPayload := plugin.GenPayload(*fetchUrl)
	if len(urlWithPayload) > 0 {
		if plugin.ParserResponse(plugin.DoRequest(urlWithPayload)) {
			storeCallback(*fetchUrl, plugin)
		}
	}
}
