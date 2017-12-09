package main

import (
	"fmt"

)
func main(){
	var x int =10//define a variable
	y:=110;
	var a[5] int//define array 01

	b:=[5] int{1,2,3,4,5}//define array o2
	const s string  = "hello world"//define a constant
	fmt.Println("aaa")//define a array
	c:=b[2:4]//数组的切片操作 b[2],b[3]不包括b[4]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a);
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(c)
	if x%2==0 {//if else
		fmt.Println("10 could be %2")
	}else{
	  fmt.Println("10 cound not be %2")
	}
	switch  x{//switch case
	case 10: fmt.Println("one")
	case 11: fmt.Println("two")
	}


}
