package main 

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:roman123456@tcp(127.0.0.1:3306)/tags")