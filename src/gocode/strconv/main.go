package main

import (
	"fmt"
	"strconv"
)

func main() {

	var strb string = "true"
	var b bool
	b, _ = strconv.ParseBool(strb)
	fmt.Printf("type:%T,value:%v\n", b, b)

	v := "43"
	var bi int64
	bi, _ = strconv.ParseInt(v, 2, 36)
	fmt.Printf("v type: %T,bi type: %T, v value:%v , bi vlaue: %v\n", v, bi, v, bi)

	var ff string = "1.222"
	var fs float64
	fs, _ = strconv.ParseFloat(ff, 10)
	fmt.Printf("type:%T ,value:%v", fs, fs)
}
