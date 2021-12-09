package main

import (
    "fmt"
    "database/sql"

   	_ "github.com/go-sql-driver/mysql"

    "Golang.com/go_mysql/db"
)

type Tag struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

func main() {
    fmt.Println("Go MySQL Tutorial")

    fmt.Println(db.Db)

    db, err := sql.Open("mysql", "root:roman123456@tcp(127.0.0.1:3306)/test")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

    if err != nil {
        panic(err.Error())
    }

    defer insert.Close()

    results, err := db.Query("select id, name from tags")

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        var tag Tag

        err = results.Scan(&tag.ID, &tag.Name)

        if err != nil {
            panic(err.Error())
        }

        fmt.Printf(tag.Name)
    }
}