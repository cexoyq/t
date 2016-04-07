package models

import (
	"fmt"
	"log"
	"crypto/md5"
	"encoding/hex"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type LoginUser struct{
	Id	int	`form:"-"`
	UserName	string
	Email string
	Level	int	`form:"-"`
	Password string	`form:"Password"`
}

func ValidateUser(username string, password string) (LoginUser, error) {
	var (
		l LoginUser
		db *sql.DB
		err  error
		//row  *sql.Row
		//rows *sql.Rows
	)
	l = LoginUser{}
	db, err = sql.Open("sqlite3", "./db.db3")
	defer db.Close()
	if err != nil {
		return l, err		//返回
	}
	md5 := md5.New()
	md5.Write([]byte(password))
	md5str := hex.EncodeToString(md5.Sum(nil))
	password = md5str	//密码转md5
	err = db.QueryRow("select id,username,email,level from user where username=? and password=? ", username, password).Scan(&l.Id, &l.UserName, &l.Email, &l.Level)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
		return l, err
	case err != nil:
		log.Fatal(err)
		return l, err
	default:
		//fmt.Printf("Username is %s\n", l.User)
		fmt.Println("sql query,id:%d,user:%s,email:%s,level:%d", l.Id, l.UserName, l.Email, l.Level)
		fmt.Println("sql query,err:", err)
		return l, nil
	}

}