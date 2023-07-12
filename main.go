package main

import (
    "net/http"

    "github.com/mstysk/todo/controller"
    "github.com/mstysk/todo/model/repository"
)

var tr = repository.NetTodo
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main(){
    server := http.Server{
        Addr: ":8080",
    }
    http.HandleFunc("/todos/", ro.HandleTodoRequest)
    server.ListenAndServe()
}
