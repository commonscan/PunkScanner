package SQLinjCrawler

import (
	"net/url"
	"strings"
)

func PayloadUrlGenerate(url url.URL) (string) {
	var q = url.Query()
	for _, i := range strings.Split(url.RawQuery, "&") {
		data := strings.Split(i, "=")
		key := data[0]
		if len(data) == 2 {
			q.Set(key, data[1]+"'")
		} else {
			return ""
		}
	}
	url.RawQuery = q.Encode()
	return url.String()
}
func GitLeakGenerate(url url.URL) (string) {
	new_url := url.Scheme + "://" + url.Host + "/.git/config"
	return new_url
}