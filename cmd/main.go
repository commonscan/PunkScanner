package main

import (
	"SQLinjCrawler/Plugin"
	"SQLinjCrawler"
	"os"
)

func main() {
	// SQLInjection scan
	//var plugin = Plugin.SQLInjectionPlugin{}
	var plugin = Plugin.S2_045{}
	SQLinjCrawler.RockIt(os.Args[1], plugin)
}
