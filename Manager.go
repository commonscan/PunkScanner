package SQLinjCrawler

import (
	"os"
	"log"
	"bufio"
	"sync"
	"fmt"
)

func Worker(conn chan string, wg *sync.WaitGroup, plugin PluginIn) {
	for ; ; {
		var rawUrl = <-conn
		Detect(rawUrl, plugin, PrintSQLInjection)
		wg.Done()
	}
}

func RockIt(path string, plugin PluginIn) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var wg = sync.WaitGroup{}
	scanner := bufio.NewScanner(file)
	var count = 0
	conn := make(chan string, 102400)
	for i := 0; i < 10000; i++ {
		go Worker(conn, &wg, plugin)
	}
	for scanner.Scan() {
		count += 1
		wg.Add(1)
		conn <- scanner.Text()
		if count%10000 == 0 {
			fmt.Println("----", count, "----")
		}
	}
	wg.Wait()
}
