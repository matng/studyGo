package main

import (
	//"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/widuu/gomysql"
	//"log"
	"encoding/json"
)

// func insert(db *sql.DB) {
// 	stmt, err := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
// 	defer stmt.Close()

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	stmt.Exec("guotie", "guotie")
// 	stmt.Exec("testuser", "123123")

// }

func main() {
	c, err := gomysql.SetConfig("/Users/Michael/dev/code/github/studyGo/src/testMysql/conf/conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	t := c.SetTable("user") //设置要处理的表名
	data := t.FindAll()     //查询表的一条数据，返回map[int]map[string]string格式

	for i, v := range data {
		xjson, err := json.Marshal(v)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("i=%v,v=%v\n", i, string(xjson))
	}
	//gomysql.Print(data)
	/*
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/uu?charset=utf8")
		if err != nil {
			log.Fatalf("Open database error: %s\n", err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		insert(db)

		rows, err := db.Query("select id, username from user where id = ?", 1)
		if err != nil {
			log.Println(err)
		}

		defer rows.Close()
		var id int
		var name string
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/
}
