package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
func main() {
	// arrMagazine := strings.Split("o l x imjaw bee khmla v o v o imjaw l khmla imjaw x", " ")
	// arrRansom := strings.Split("imjaw l khmla x imjaw o l l o khmla v bee o o imjaw imjaw o", " ")
	arrMagazine := strings.Split("give me one grand today night", " ")
	arrRansom := strings.Split("give night me one grand today", " ")

	var magMap map[uint64]int
	magMap = make(map[uint64]int)
	for _, t := range arrMagazine {
		if _, ok := magMap[hash(t)]; ok == true {
			magMap[hash(t)] = magMap[hash(t)] + 1
		}else   {
            magMap[hash(t)] = 1
        }
		
	}
	result := "YES"
	for i := 0; i < len(arrRansom); i++ {
		_, ok := magMap[hash(arrRansom[i])]
		if ok == false || (ok == true && magMap[hash(arrRansom[i])] < 1) {
			result = "NO"
			break
		} else {
			magMap[hash(arrRansom[i])] = magMap[hash(arrRansom[i])] - 1

		}
	}
	fmt.Println(result)

}
