package repository

import (
    "database/sql"
    "fmt"
)

var Db *sql.DB

func intt() {
    var err error
    dataSourceName := fmt.Sprint("%s:%s@tcp(%s)/%s?charset=utf8",
        "todo-app", "todo-password", "sample-api-db:3306", "todo",
    )
    Db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        panic(err)
    }
}
