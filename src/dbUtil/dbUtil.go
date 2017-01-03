package dbUtil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Insert(userName string, password string) {

	db, err := sql.Open("mysql", "root:123456@/uu")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec(userName, password)
}
