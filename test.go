package main

import (
  "database/sql"
   _ "github.com/mattn/go-sqlite3"
  // "net/http"
  // "html/template"
  // "strconv"
  "fmt"
)

type User struct {
  Id int64
  Name string
  Balance int16
  Status int16
}


func main() {

  db, err := sql.Open("sqlite3", "users.db")
  if err != nil {
    panic(err)
  }

  defer db.Close()
  row := db.QueryRow("select * from Users where id = $1", 1)
  if err != nil {
    panic(err)
    }

  current_user := User{}
  err = row.Scan(&current_user.Id, &current_user.Name, &current_user.Balance, &current_user.Status)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(current_user.Id, current_user.Name, current_user.Balance, current_user.Status)
}
