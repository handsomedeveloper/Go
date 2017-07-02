package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:SheJie_123456@/test?charset=utf8")
	defer db.Close()
	checkErr(err)

	//insert data
	stmt, err := db.Prepare("INSERT INTO user_info (username, departname, created) VALUE (?, ?, ?)")
	checkErr(err)
	defer stmt.Close()
	res, err := stmt.Exec("SheJie", "IT部", "2017-6-15")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//update data
	statementUpdate, err := db.Prepare("UPDATE user_info SET username = ? WHERE uid = ?")
	checkErr(err)
	defer statementUpdate.Close()
	resultUpdate, err := statementUpdate.Exec("SheUpdate", id)
	checkErr(err)
	affect, err := resultUpdate.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//read data
	rows, err := db.Query("SELECT * FROM user_info WHERE 1 = ?", 1)
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var uid int
		var username string
		var departName string
		var created string
		err := rows.Scan(&uid, &username, &departName, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departName)
		fmt.Println(created)
	}

	//delete data
	statementDelete, err := db.Prepare("DELETE FROM user_info WHERE uid = ?")
	checkErr(err)
	defer statementDelete.Close()
	resultDelete, err := statementDelete.Exec(id)
	checkErr(err)
	affectDelete, err := resultDelete.RowsAffected()
	checkErr(err)
	fmt.Println(affectDelete)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
