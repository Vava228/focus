package main

import (
  "net/http"
  "html/template"
  "strconv"
  // "fmt"
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
  http.HandleFunc("/user/", user)
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
  bob := User{Id: 12, Name: "Bob", Balance: 100, Status: 0}
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
      if err != nil || id < 1 {
          http.NotFound(w, r)
          return
        }
    files := []string {
    "templates/user.html",
    "templates/base.html",
    }
    tmpl, _ := template.ParseFiles(files...)
    tmpl.Execute(w, bob)
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
