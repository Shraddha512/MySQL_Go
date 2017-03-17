package main

//database/sql --> sql interface
//github.com/go-sql-driver/mysql --> driver, supporting Go sql interface
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//sql.Open(name, data source name)
	//DSN format: username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", "foo:foofoo@127.0.0.1:3306/test?charset=utf8")
	checkErr(err)
	fmt.Println("Registered database")

	var name = "userinfo"
	//create
	/*
		_, err = db.Exec("CREATE DATABASE " + name)
		checkErr(err)
		fmt.Println("Created database")	*/

	//insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "xyz", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("axrsbxs", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//query
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
	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
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

/*

eg. TABLE

CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);*/
