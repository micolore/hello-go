// dbutils
package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:123Moppo!@#@(192.168.1.5:3306)/around?charset=utf8")
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("mysql collenction error")
		return
	}
	fmt.Println("mysql collection success!")

	/*
		stmt, err := db.Prepare("INSERT client_version SET content=?")
		checkErr(err)

		res, err := stmt.Exec("test mysql")
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Printf("insert id %d\n", id)
	*/

	//如果不加要查询的字段就不会取到值
	rows, _ := db.Query("select id,content from client_version")
	var array []string
	//打印返回的列字段
	array, err = rows.Columns()
	fmt.Println(array)

	for rows.Next() {
		var id uint64
		var content string
		err = rows.Scan(&id, &content)
		fmt.Println(id, "---", content)
	}
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
