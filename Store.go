package SQLinjCrawler

import (
	"database/sql"
	"fmt"
)

func StoreToMysql(url string) {
	db, err := sql.Open("mysql", "root:xxx@tcp(178.128.67.122:3306)/sqlinj")
	if err != nil {
		print(err)
		panic("failed to connect to mysql")
	}
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO url (url) VALUES( ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(url)
}

func PrintSQLInjection(url string) {
	fmt.Println("SQL Inj:", url)
}
