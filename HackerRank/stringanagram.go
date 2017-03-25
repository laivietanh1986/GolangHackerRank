package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"c", "d", "e"}
	b := []string{"a", "b", "c"}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				a[i] = "0"
				b[j] = "0"
				break
			}
		}
	}
	chuoi := "VietAnh"
	fmt.Println(CountRemainChar(a) + CountRemainChar(b))
	fmt.Println(strings.Split(chuoi, ""))
	for t, a := range strings.Split(chuoi, "") {
		fmt.Println(t, a)
	}

}
func CountRemainChar(chuoi []string) int {
	count := 0
	for index := 0; index < len(chuoi); index++ {
		if chuoi[index] != "0" {
			count++
		}
	}
	return count
}
