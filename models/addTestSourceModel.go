package models

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
type Source struct{
	Id			int `json:"id"`
	Name 		string `json:"name"`
	Rtmp 		string `json:"rtmp"`
	ParentId	int	`json:"parentid"`
}
type Parent struct{
	Id			int
	Name		string
	Sources		[]Source
}

func RetuTestSource()([]Parent,error){
	var(
		//S		[]Source
		P 		[]Parent
		err 	error
		db		*sql.DB
		sqlstr1	string
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
	//查询父级列表
	sqlstr=`select id,name from test where parentid=0`
	rows, err = db.Query(sqlstr)
	if err != nil {
	    log.Printf("db.Query() Err :",err)
		return P,err
	}
	defer rows.Close()
	i:=0
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
	rows.Close()
	/**上面已取到数据库中的父级了**/
	/**查询父级下面的数据**/
	sqlstr1 = `select count(*) from test where parentid=?`
	sqlstr = `select id,name,rtmp from test where parentid=?`
	i=0;
	for _,v := range P {
		//先取得数据查询后的总行数
		var (
			count int64
		)
		err = db.QueryRow(sqlstr1,v.Id).Scan(&count);
		if(err != nil){break;}
		log.Printf("Parnent  Id:%d, Name:%s",v.Id,v.Name)
		log.Print("sqlstr:",sqlstr)
		log.Printf("rows count:",count)
		if(err != nil){
			continue;
		}
		P[i].Sources = make([]Source,count);
		v.Sources = make([]Source,count);
		if count == 0 {continue;}
		//查询二级
		rows,err = db.Query(sqlstr,v.Id);
		defer rows.Close()
		j:=0;
		for rows.Next(){
			err = rows.Scan(&P[i].Sources[j].Id,&P[i].Sources[j].Name,&P[i].Sources[j].Rtmp);
			//err = rows.Scan(&v.Sources[i].Id,&v.Sources[i].Name,&v.Sources[i].Rtmp)
			if err != nil {
				break;
			}
			log.Printf("child ID:%d,Name:%s,Rtmp:%s",P[i].Sources[j].Id,P[i].Sources[j].Name,P[i].Sources[j].Rtmp)
			//log.Printf("child ID:%d,Name:%s,Rtmp:%s",v.Sources[i].Id,v.Sources[i].Name,v.Sources[i].Rtmp)
			j++;
		}
		rows.Close()
		//num := copy(P[i].Sources,v.Sources);
		//log.Printf("copy slice num:",num)
		i++;
	}
	println("data:",P)
	return P,err
}

func addTestSource(data Source) error{
	var(
		db 			*sql.DB
		ret 		sql.Result
		sqlstr		string
		err 		error;
	)
	db, err = sql.Open("sqlite3", "./db.db3")
	defer db.Close()
	if err != nil {
		log.Printf("call RetuTestSource() err :",err)
		return  err		//返回
	}
	sqlstr=`insert into test (name,rtmp,parentid) values(?,?,?)`
	ret,err = db.Exec(sqlstr,data.Name,data.Rtmp,data.ParentId);
	if(err!=nil){
		return err;
	}
	count,err := ret.RowsAffected();
	lastid,err := ret.LastInsertId();
	log.Printf("insert row count:%d, id:%d",count,lastid)
	return err;
}