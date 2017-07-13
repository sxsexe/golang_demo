package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
	db, e := sql.Open("mysql", "root:sxsexe@tcp(localhost:3306)/mysql?charset=utf8")
	if e != nil {
		fmt.Println("e : ", e)
		return
	}

	_, e2 := db.Query("select 1")
	if e2 != nil {
		fmt.Println("e2 : ", e2)
	} else {
		println("DB OK")
		rows, err3 := db.Query("select password from user")
		//		println(rows, ", ", err3)
		fmt.Println(err3)

		if rows != nil {
			for rows.Next() {
				var name string
				var password string
				row_err := rows.Scan(&name, &password)
				if row_err != nil {
					println("row_err = ", row_err)
				} else {
					println("username = ", name, ", password = ", password)
				}
			}
		} else {
			println("Rows is empty")
		}
	}
}
