package main

import (
	"SQLinjCrawler/Plugin"
	"SQLinjCrawler"
)

func main() {
	// SQLInjection scan
	//var plugin = Plugin.SQLInjectionPlugin{}
	//SQLinjCrawler.RockIt("/tmp/1.txt", plugin)
	var plugin = Plugin.S2_045{}
	SQLinjCrawler.RockIt("/tmp/2.txt", plugin)
}
