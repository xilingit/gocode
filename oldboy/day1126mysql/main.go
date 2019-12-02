package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

/*
	下载依赖 go get -u github.com/go-sql-driver/mysql

	普通SQL语句执行过程
		1、客户端对SQL语句进行占位符替换得到完整的SQL语句
		2、客户端发送完整SQL语句到mysql服务端
		3、mysql服务端执行完整的SQL语句并将结果返回给客户端

	预处理执行过程
		1、把SQL语句分成两部分，命令部分与数据部分
		2、先把命令部分发送给mysql服务端，mysql服务端进行SQL预处理
		3、然后把数据部分发送给mysql服务端，mysql服务端对SQL语句进行占位符替换
		4、mysql服务端执行完整的SQL语句并将结果返回给客户端

	为什么要预处理
		1、优化mysql服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本
		2、避免SQL注入问题

	Go实现MySQL预处理
	Prepare(query string)
	Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令，返回值可以同时执行多个查询和命令

	Go语言实现事务
		事务：一个最小的不可再分的工作单元，通常一个事务对应一个完整的业务，同时这个完整的业务需要执行多次
		的DML语句共同联合完成，A转账给B，这里面就需要执行两次update操作
		MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务，事务处理可以用来维护数据库的完整性，
		保证成批的SQL语句要么全部执行，要么全部不执行
	
	事务的ACID
		通常事务必须满足4个条件：原子性、一致性、
*/
func main() {
	//使用mysql驱动
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	checkErr(err)
	defer db.Close()
	//SetMaxOpenConns:设置与数据库建立连接的最大数目
	//SetMaxIdleConns:设置连接池中的最大闲置连接数
	err = db.Ping() //尝试与数据库建立连接
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	/*单行查询
	db.QueryRow()执行一次查询
	*/
	// sqlStr := "select * from stuinfo where id=?"
	// var user User
	// //非常重要，要确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	// err = db.QueryRow(sqlStr, 3).Scan(&user.id, &user.name)
	// if err != nil {
	// 	fmt.Println("scan failed,err:", err)
	// 	return
	// }
	// fmt.Printf("id:%d,name:%s\n", user.id, user.name)

	//查询多条数据
	// sqlStr := "select * from stuinfo"
	// rows, err := db.Query(sqlStr)
	// if err != nil {
	// 	fmt.Println("query err,err:", err)
	// 	return
	// }
	// defer rows.Close()
	// //循环读取结果中的数据
	// for rows.Next() {
	// 	var u User
	// 	err := rows.Scan(&u.id, &u.name)
	// 	if err != nil {
	// 		fmt.Println("scan err,err:", err)
	// 		return
	// 	}
	// 	fmt.Printf("id:%d,name:%s\n", u.id, u.name)
	// }

	//插入数据
	// sqlStr := "insert into stuinfo(id,name) values (?,?)"
	// result, err := db.Exec(sqlStr, 2, "hello")
	// if err != nil {
	// 	fmt.Println("insert failed,err:", err)
	// 	return
	// }
	// id, err := result.LastInsertId() //新插入数据的id
	// if err != nil {
	// 	fmt.Println("get lastinsert id failed,err:", err)
	// 	return
	// }
	// fmt.Printf("insert success,the id is %d\n", id)

	//更新数据
	// sqlStr := "update stuinfo set name=? where id = ?"
	// result, err := db.Exec(sqlStr, "world", 2)
	// if err != nil {
	// 	fmt.Printf("update failed,err:%v\n", err)
	// 	return
	// }
	// n, err := result.RowsAffected() //操作影响的行数
	// if err != nil {
	// 	fmt.Printf("get RowsAffected failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("update success,affected rows:%d\n", n)

	//删除数据
	// sqlStr := "delete from stuinfo where id = ?"
	// result, err := db.Exec(sqlStr, 3)
	// if err != nil {
	// 	fmt.Printf("delete failed,err:%v\n", err)
	// 	return
	// }
	// n, err := result.RowsAffected() //操作影响的行数
	// if err != nil {
	// 	fmt.Printf("get RowsAffected failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("delete success,affected rows:%d\n", n)

	//预处理查询示例
	sqlStr := "select id,name from stuinfo where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Printf("query failed,err%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d,name:%s\n", u.id, u.name)
	}
}

//User 结构体
type User struct {
	id   int
	name string
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}
