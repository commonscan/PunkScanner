package SQLinjCrawler

import (
	"testing"
)

func TestPayloadUrlGenerate(t *testing.T) {
}

func TestCrawl(t *testing.T) {
	doRequest("http://www.baidu.com")
}

func TestDetectSQLInj(t *testing.T) {
	//Detect("http://140.82.4.59:8080/sqlmap/mysql/get_int_user.php?id=1", PayloadUrlGenerate, SQLInjParser, PrintSQLInjection)
}

func TestRockIt(t *testing.T) {
	//RockIt("./urls")
}

func TestStoreSQLInjection(t *testing.T) {
	StoreToMysql("http://www.baidu.com")
}
