package main

import (
  "database/sql"
   _ "github.com/mattn/go-sqlite3"
  "net/http"
  "html/template"
  "strconv"
  "fmt"
)

type User struct {
  Id int64
  Name string
  Balance int16
  Status int16
}

type Tool struct {
  Id int64
  Name string
  Status string
  Price float32
}

func Greetings(word string) (text string) {
  text = "You`re on the " + word + "."
  return text
}

func main() {

  HandlerSettings()

}

func HandlerSettings() {
  // handlers
  http.HandleFunc("/", index)
  http.HandleFunc("/price/", price)
  http.HandleFunc("/contact/", contact)
  http.HandleFunc("/user", user)
  //local server creatinng and setting
  http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
  bob := User{Id: 1, Name: "Bob", Balance: 100, Status: 0}

  files := []string {
    "templates/index.html",
    "templates/base.html",
  }

  tmpl, _ := template.ParseFiles(files...)
  tmpl.Execute(w, bob)
}

func contact(w http.ResponseWriter, r *http.Request) {
  bob := User{Id: 1, Name: "Bob", Balance: 100, Status: 0}

  files := []string {
    "templates/contact.html",
    "templates/base.html",
  }

  tmpl, _ := template.ParseFiles(files...)
  tmpl.Execute(w, bob)

}

func user(w http.ResponseWriter, r *http.Request) {

  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  files := []string {
    "templates/user.html",
    "templates/base.html",
  }

  db, err := sql.Open("sqlite3", "users.db")
  if err != nil {
    panic(err)
  }

  row := db.QueryRow("select * from Users where id = $1", id)
  if err != nil {
    panic(err)
  }

  current_user := User{}
  err = row.Scan(&current_user.Id, &current_user.Name, &current_user.Balance, &current_user.Status)
  if err != nil {
    fmt.Println(err)
  }

  defer db.Close()

  tmpl, err := template.ParseFiles(files...)
  if err != nil {
    http.NotFound(w, r)
  }
  tmpl.Execute(w, current_user)
}

func price(w http.ResponseWriter, r *http.Request) {
  bob := User{Id: 1, Name: "Bob", Balance: 100, Status: 0}

  files := []string {
    "templates/price.html",
    "templates/base.html",
  }

  tmpl, _ := template.ParseFiles(files...)
  tmpl.Execute(w, bob)

}
