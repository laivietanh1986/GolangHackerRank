package Todoes
import (
    "html/template"
    "net/http"
)
type Todo struct {
    Task string
    Done bool
}
func main()  {
    tmpl:= template.Must(template.ParseFiles("hello/todos.html"))
    todos:= []Todo{
        {"Learn Go",true},
        {"Create Website For Uncle",false},
        {"cycling at saturday",true},
    }
    http.HandleFunc("/",func(w http.ResponseWriter , r*http.Request){
        tmpl.Execute(w,struct{Todos []Todo}{todos})
    })
    http.ListenAndServe(":8080",nil)
}