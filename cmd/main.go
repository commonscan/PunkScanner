package main

import (
	"SQLinjCrawler/Plugin"
	"SQLinjCrawler"
)

func main() {
	// SQLInjection scan
	//var plugin = Plugin.SQLInjectionPlugin{}
	//SQLinjCrawler.RockIt("/tmp/1.txt", plugin)
	var leakPlugin = Plugin.GitLeakPlugin{}
	SQLinjCrawler.RockIt("/tmp/2.txt", leakPlugin)
}
