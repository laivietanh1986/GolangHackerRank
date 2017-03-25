package main
import (
    "net/http"
)
func main()  {
    fs:=http.FileServer(http.Dir("templates/"))
    http.Handle("/static/",http.StripPrefix("/static",fs))
    http.ListenAndServe(":8080",nil)
}