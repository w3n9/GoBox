package db

import (
	"GoBox/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)
var Conn *sql.DB
var err error
func init(){
	config:=config.Get()
	Conn,err=sql.Open(config.DataSource.Type,config.DataSource.Url)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

