package main

import "fmt"

func main()  {

	var x int8 =  127// -128 --- 127
	var y int16 = 32767 // (-128 - 3267)
	//var z int32 = 1.8446744e+19
	fmt.Println(x)
	fmt.Println(y)
	//fmt.Println(z)

	//Using 8-bit unsigned int
	var xx uint8 = 225
	fmt.Println(xx)

	var yy int16 = 32767
	fmt.Println(yy+2, yy-2)


}
// 65535