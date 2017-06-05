package main


import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {

	db, err := sql.Open("mysql", "ayong:zzz@/test?charset=utf8")
	checkErr(err)

	//------------------------------------------------------------------------------
	// insert
	fmt.Println("aa1")
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	res, err := stmt.Exec("HELLO_abc", "MARKETING", "2012-12-09")
	if res==nil {fmt.Println("nil!!!!")}
	//id, err := res.LastInsertId()
	//fmt.Println(id)
	fmt.Println("aa2")

	//checkErr(err)
	/*
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("NATEYONG", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	*/

	/*
	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
*/

	/*
	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	*/

	db.Close()

}

func checkErr(err error) {
	//fmt.Println("Ok. no error found")
	if err != nil {
		fmt.Println("Found an error!")
		panic(err)
	}
}
