package controller

import (
    "net/http"
)

type Router interface {
    HandleTodoRequest(w http.ResponseWritter, r *http.Request)
},

type router struct {
    tc TodoContoroller
}

func NewRouter(tc TodoContoller) Router {
    return &router{tc}
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
    switch r.Method{
        case "GET":
            ro.tc.GetTodos(w, r)
        case "POST":
            ro.tc.PostTodo(w, r)
        case "PUT":
            ro.tc.PutTodo(w, r)
        case "DELETE":
            ro.tc.Delete(w, r)
        default:
            w.WriteHeader(405)
    }
}
