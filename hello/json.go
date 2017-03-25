package main
import (
    "fmt"
    "net/http"
    "encoding/json"
)
type User struct {
    FirstName string `json:"firstname"`
    LastName string  `json:"lastname"`
    Age int `json:"age"`
}
func main()  {
    http.HandleFunc("/decode",func (w http.ResponseWriter , r*http.Request)  {
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        fmt.Fprintf(w,"%s %s is %d \n",user.FirstName,user.LastName,user.Age)
    })
    http.HandleFunc("/encode",func (w http.ResponseWriter , r*http.Request)  {
        peter:=User{
            FirstName:"Lai",
            LastName:"Viet Anh",
            Age:31,
        }
        json.NewEncoder(w).Encode(&peter)
    })
    http.ListenAndServe(":8080",nil)

}

