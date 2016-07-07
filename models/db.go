package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const (
	DB_HOST = "tcp(192.168.0.70:3306)"
	DB_NAME = "nava"
	DB_USER = /*"root"*/ "root"
	DB_PASS = /*""*/ "expert"
)

func init() {
	// register model
	orm.RegisterModel(new(User), new(Profile))
	// set default database
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn, 30)
	fmt.Println("Registered Database")

	o := orm.NewOrm()
	user := User{Username: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Username = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

