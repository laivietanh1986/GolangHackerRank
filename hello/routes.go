package main
import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)
func main()  {
    userAges := map[string]int{
        "Alice":30,
        "Bob":25,
        "Clarie":27,
    }
    r:=mux.NewRouter()
    r.HandleFunc("/user/{name}",func(w http.ResponseWriter , r * http.Request)  {
        vars := mux.Vars(r)
        name:= vars["name"]
        age:= userAges[name]
        fmt.Fprintf(w,"%s is %d years old \n",name,age)

    }).Methods("GET")
    http.ListenAndServe(":8080",r)
}