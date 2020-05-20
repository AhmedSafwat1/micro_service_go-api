package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
const (
	mysql_user_name     = "mysql_user_name"
	mysql_user_password ="mysql_user_password"
	mysql_user_host		= "mysql_user_host"
	mysql_user_shema	= "mysql_user_schema"
)

var (
	Client *string
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"",
		"127.0.0.1:3306",
		"users_db",
	)
	g ,err := sql.Open("mysql", dataSourceName)
	fmt.Println(g)
	x := "Dsdds"
	Client = &x

	if err != nil {
		panic(err)
	}
	//if err = Client.Ping(); err != nil {
	//	panic(err)

	//}
	log.Println("database connect  from init hi ..... ")

}

func Connect() *sql.DB {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			"root",
			"",
			"127.0.0.1:3306",
			"users_db",
		)
	Client ,err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connect hi ..... ")
	return Client;
}