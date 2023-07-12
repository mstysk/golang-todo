package controller

import (
    "encoding/json"
    "net/http"
    "path"
    "strconv"

    "github.com/mstysk/todo/controller/dto"
    "github.com/mstysk/todo/model/entity"
    "github.com/mstysk/todo/model/repository"
)

type TodoController interface {
    GetTodos(w http.ResponseWriter, r *http.Request)
    PostTodo(w http.ResponseWriter, r *http.Request)
    PutTOdo(w http.ResponseWriter, r *http.Request)
    DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
    tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
    return &todoController{todo}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := tc.tr.GetTodos()
    if err != nil {
        w.WriteHeader(500)
        return
    }

    var todoResponses []dto.todoResponse
    for _, v := range todos {
        todoResponse = append(todoResponse, dto.TodoResponse{Id: v.id, Title: v.title, Content: v.content})
    }

    var todoResponse dto.todosResponse
    todoRespnse.Todos = todoResposes

    output, _ := json.MarshalIndent(todoResponse.Todos, "", "\t\t")


    w.Header().Set("Content-Type", "applicatoin/json")
    w.Writer(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r* http.Request) {
    body := make([]byte, r.ContentLength)
    r.Body.Read(body)
    var todoRequest dto.todoRequest
    json.Unmarshal(body, &todoRequest)

    todo := entity.TOdoEntity{Title: todoRequest.Title, Content: todoRequest.Content}

    id, err := tc.tr.InsertTodo(todo)
    if err != nil {
        w.WriteHeader(500)
        return
    }

    w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
    w.WriteHeader(201)
}

func (tc *todoController) PutTodo(w http.ResponseWriter, r* http.Request) {
    todoId, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        w.WriterHeader(400)
        return
    }

    body := make([]byte, r.ContentLength)
    r.Body.Read(body)
    bar todoRequest dto.todoRequest
    json.Unmarshal(body, &todoRequest)

    todo := entity.TodoEntity{Id: todoId, Title: todoRequest.Title, Content: todoRequest.Content}

    err = tc.tr.UpdateTodo(todo)

    if err != nil {
        w.WeriteHeader(500)
        return
    }

    w.WriterHeader(204)
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r* http.Request) {
    todoId, err := strconv.Atoi(path)
    if err != nil {
        w.WriterHeader(400)
        return
    }

    err = tc.tr.DeleteTodo(todoId)
    if err != nil {
        w.WriterHeader(500)
        return
    }

    w.WriteHeader(204)
}
