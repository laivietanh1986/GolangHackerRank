package main
import (
    "net/http"
    "html/template"
)
type ContactDetail struct {
    Email string    
    Subject string
    Message string
}
func main()  {
    tmpl:=template.Must(template.ParseFiles("templates/Form.html"))
    http.HandleFunc("/",func (w http.ResponseWriter,r * http.Request)  {
        if r.Method != http.MethodPost {
           tmpl.Execute(w,nil)  
           return
        }
        detail:= ContactDetail{
            Email:r.FormValue("Email"),
            Subject:r.FormValue("Subject"),
            Message:r.FormValue("Message"),
        }
        _=detail
        tmpl.Execute(w,struct{Success bool}{true})
    })
    http.ListenAndServe(":8080",nil)
}