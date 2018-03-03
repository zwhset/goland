package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/* version1 */
	//s, sep := "", ""
	//var s, sep string
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//
	//}
	// fmt.Println(s)

	// version2 怎么消耗时间反而长了？
	fmt.Println(strings.Join(os.Args[1:], " "))
}
