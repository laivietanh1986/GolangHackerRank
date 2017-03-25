package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
   var length,numrot int
   io := bufio.NewReader(os.Stdin)
   fmt.Fscan(io, &length)
   fmt.Fscan(io,&numrot)
   arr:= make([]int,length)
   for index :=0;index <length;index ++{
       fmt.Fscan(io,&arr[index])        
   }
  
    num:=numrot
    if numrot >length {
        num = numrot -length
    }
    num = num % length
    result := rotation(arr,length,num)
    for index:=0;index<length;index++{
        fmt.Printf("%v ",result[index])
    }
    
    
}
func rotation(arr []int,len int,num int) []int {
    temp:=make([]int,len)

    for index := 0; index < len ; index++ {
        
        position:=index-num
        if position< 0 {
            temp[len+position]=arr[index]    
        }else {
            temp[position] = arr[index]
        }
        
        
    }
    return temp
}