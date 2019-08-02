package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id  		int      		// `orm:pk, auto`
	Name 		string   		`orm:size(100)`
	order 		[]*User_order	`orm:ref(fk)`
}

type User_order struct {
	Id 			int
	Order_data 	string	`orm:size(100)`
	user        *User	`orm:"reverse(many)"`
}


//	create database class27;

//	create table user{....}

func init()  {
	// set default databases
	orm.RegisterDataBase("default","mysql","root@itcast@tcp(127.0.0.1:3306)/class27?charset=utf-8", 30)

	// register model
	orm.RegisterModel(new(User), new(User_order))

	// create table
	orm.RunSyncdb("default", true, true)
}