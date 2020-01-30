package main

import (
	"fmt"
	"runtime"
)


func Add(x,y int) int {
	z :=x+y;
	fmt.Printf(".>  %d !*\n" ,z)
	runtime.Gosched();//让出时间片
	return z
}

func main(){
	f :=-4;
	k := f;
	for i:=0;i<10;i++{
		go Add(i,i);
		//fmt.Printf(".> \n %d !*" ,c)
	}
	runtime.Gosched();
	fmt.Printf(".>  %d !*\n" ,k)



}
