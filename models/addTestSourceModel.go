package models

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
type Source struct{
	Id			int `json:"id"`
	Name 		string `json:"name"`
	Rtmp 		*string `json:"rtmp"`
}
type Parent struct{
	Id			int
	Name		string
	Sources		[]Source
}
func RetuTestSource()([]Parent,error){
	var(
		P 		[]Parent
		//S 		[]Source
		err 	error
		db		*sql.DB
		sqlstr	string
		rows	*sql.Rows
	)
	log.Printf("call RetuTestSource()")
	db, err = sql.Open("sqlite3", "./db.db3")
	defer db.Close()
	if err != nil {
		log.Printf("call RetuTestSource() err :",err)
		return P, err		//返回
	}
	//查询二级列表中父级的总数
	sqlstr=`select count(*) from test where parentid=0`
	var parentCount int
	err = db.QueryRow(sqlstr).Scan(&parentCount)
	if(err != nil){
		log.Printf("db.QueryRow,Parent Count Err:",err)
	}
	log.Printf("db.QueryRow,Parent Count :",parentCount)
	P=make([]Parent,parentCount,parentCount)	//初始化父级切片
	//S=make([]Source,parentCount,parentCount)	
	//查询父级列表
	sqlstr=`select id,name from test where parentid=0`
	rows, err = db.Query(sqlstr)
	if err != nil {
	    log.Printf("db.Query() Err :",err)
		return P,err
	}
	defer rows.Close()
	var i int
	i=0
	for rows.Next() {
		var(
			//id 		int
			//name 	string
			//rtmp	string
		)
	    //if err := rows.Scan(&id,&name); err != nil {
		if err := rows.Scan(&P[i].Id,&P[i].Name); err != nil {
	        log.Printf("rows.Scan Err:",err)
	    }
	    log.Printf("id:%d,name:%s",P[i].Id,P[i].Name)
		i++
	}
	if err := rows.Err(); err != nil {
	    log.Fatal(err)
	}
	return P,err
}