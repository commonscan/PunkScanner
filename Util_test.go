package SQLinjCrawler

import (
	"testing"
	"net/url"
)

func TestPayloadUrlGenerate(t *testing.T) {
	testUrl, err := url.Parse("http://localhost/vulnerabilities/sqli/?id=1&Submit=Submit")
	if err != nil {
		panic(err)
	}
	PayloadUrlGenerate(*testUrl)
}

func TestCrawl(t *testing.T) {
	doRequest("http://www.baidu.com")
}

func TestDetectSQLInj(t *testing.T) {
	Detect("http://140.82.4.59:8080/sqlmap/mysql/get_int_user.php?id=1", PayloadUrlGenerate, SQLInjParser, PrintSQLInjection)
}

func TestRockIt(t *testing.T) {
	RockIt("./urls")
}

func TestStoreSQLInjection(t *testing.T) {
	StoreToMysql("http://www.baidu.com")
}
