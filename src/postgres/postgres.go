package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=SheJie password=SheJie123456 dbname=mydb sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("insert into user_info(username,department,created) values ($1,$2,$3) returning uid")
	_, err = stmt.Exec("SheJie", "研发部", "2017-07-02")
	checkErr(err)

	var insertId int
	err = db.QueryRow("insert into user_info(username,department,created) values ($1,$2,$3) returning uid", "SheJie", "研发部", "2017-07-02").Scan(&insertId)
	checkErr(err)
	fmt.Println("insertId: ", insertId)

	stmt, err = db.Prepare("update user_info set username=$1 where uid=$2")
	checkErr(err)
	res, err := stmt.Exec("MrJie", 1)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	rows, err := db.Query("select uid,username,department,created from user_info")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	stmt, err = db.Prepare("delete from user_info where uid=$1")
	checkErr(err)
	res, err = stmt.Exec(insertId)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
