package Plugin

import (
	"net/url"
	"github.com/valyala/fasthttp"
	"regexp"
	"fmt"
	"strings"
)

type GitLeakPlugin struct {
	name string
}

func (plugin GitLeakPlugin) GenPayload(url url.URL) string {
	new_url := url.Scheme + "://" + url.Host + "/.git/config"
	return new_url
}

func (plugin GitLeakPlugin) ParserResponse(response *fasthttp.Response) bool {
	if response.Header.StatusCode() == 200 {
		str := response.Body()
		var errorRegexp, _ = regexp.Compile("core|branch|remote")
		if errorRegexp.Match(str) && !strings.Contains(string(str), "</head>") {
			return true
		}
	}
	return false
}
func (plugin GitLeakPlugin) GetName() string {
	return "GITLEAK"
}
func (plugin GitLeakPlugin) GenInfo(url url.URL) string {
	return fmt.Sprintf("GITLEAK %s", plugin.GenPayload(url))
}
