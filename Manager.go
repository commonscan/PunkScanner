package SQLinjCrawler

import (
	"os"
	"log"
	"bufio"
	"sync"
	"fmt"
)

func Worker(conn chan string) {
	for ; ; {
		var raw_url = <-conn
		Detect(raw_url, PayloadUrlGenerate, SQLInjParser, StoreToMysql)
	}
}

func RockIt(path string) {
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
		go Worker(conn)
	}
	// 8750000
	for scanner.Scan() {
		count += 1
		if count < 8750000 {
			scanner.Text()
			continue
		}
		conn <- scanner.Text()
		if count%10000 == 0 {
			fmt.Println("----", count, "----")
		}
	}
	fmt.Println("start waiting")
	wg.Wait()
	fmt.Println("end")
}
