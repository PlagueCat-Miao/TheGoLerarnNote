package main

import (
	"fmt"
	"runtime"
	"time"
)


func Add(x,y int) int {
	z :=x+y;
	fmt.Printf(".>  %d !*\n" ,z)
	runtime.Gosched();//让出时间片
	return z
}
/*

	f :=-4;
	k := f;
	for i:=0;i<10;i++{
		go Add(i,i);
		//fmt.Printf(".> \n %d !*" ,c)
	}
	runtime.Gosched();
	fmt.Printf(".>  %d !*\n" ,k)

*/

func test(c chan int){
	c<-'A'
}

func testDeadLock(c chan int){
	for{
		fmt.Println(<-c)
	}
}

func main(){
	var miao chan int
	miao =make(chan int,1024);
	//如果channel不带1024

	miao <- 1;
	fmt.Printf(".>  input %d\n" ,1)
	kk,x := <-miao
	fmt.Printf(".>  output %d\n" ,kk)

	if(x==true){fmt.Printf("miao runing\n" )}
	close(miao);

	_,xx := <-miao

	fmt.Printf("main close" ,1)
	if(xx==true){fmt.Printf("miao runing" ,1)}


		c :=make(chan int)
		go test(c)
		go testDeadLock(c)
		time.Sleep(time.Millisecond)


}
