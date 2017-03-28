package main
import (
    "fmt"
    "strings"
)
func main()  {
    fmt.Println("Balance Bracket")
    CheckIsBalance("{[()]}")
    CheckIsBalance("{[(])}")
    CheckIsBalance("{{[[(())]]}}")
    

}
func CheckIsBalance(s string) {
    var arr[]string 
    openBrackets:=map[string]bool{"{":true,"[":true,"(":true}
    closeBrackets:=map[string]string{"}":"{","]":"[",")":"("}
    arrchars:=strings.Split(s,"")   
    for _,char := range arrchars {
        if _,ok:=openBrackets[char];ok == true{
            arr = append(arr,char)
            fmt.Println(arr)
        }else if openchar,ok:=closeBrackets[char];ok == true{                      
            if arr[len(arr)-1] == openchar {
                arr = arr[:len(arr)-1]    
            }else {    
                fmt.Println("No")
                return
            }
        }
    }
    if len(arr) == 0 {
        fmt.Println("Yes")
    }else  {
        fmt.Println("No")
    }

}