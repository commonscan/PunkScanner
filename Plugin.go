package SQLinjCrawler

import (
	"net/url"
	"github.com/valyala/fasthttp"
)

type PluginIn interface {
	GetName() (string)
	GenPayload(url url.URL) (string)
	ParserResponse(response *fasthttp.Response) (bool)
	GenInfo(url url.URL) (string)
	DoRequest(url string) *fasthttp.Response
}
