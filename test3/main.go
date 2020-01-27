package main

import (
	"fmt"

)

import toolNya2 "github.com/TheGoLerarnNote/test3/toolNya"

func exp(x int)(int ){ //形参
	if  x=3;x== 5 {
		return 7;
	}else {
		return x-1;
	}

	//return -7;
}

func main() {
	 var y int =5;
	fmt.Printf("%d %d",exp(y),y);
	switch  {
	case 0<=y && 5 >= y:
		fmt.Printf(" 0~5");
	case 5<=y && 10 >= y:
		fmt.Printf(" 5~10");
	}
    a := []int {1,2,3,4,5}//这是切片
    maopao:
    for i,j := 0, len(a);i<j;i++{

    	for k := 0;k<j-i-1;k++{
			if i==2 {
				//break maopao; //多级跳,相当于直接到maopao级 写了break;
				continue maopao;
			}
    	    if a[k]<a[k+1]{
    	    	a[k+1],a[k]=a[k],a[k+1];

			}

		}

	}
	for i,j := 0, len(a);i<j;i++{
		fmt.Printf("%d\n",a[i]);
	}

	//[]int 为数组切片 其性质接近vector或数组 ，数组切边名之间的赋值属于指针操作，结果两个名变为同名
	//[6]int 为数组 ，其性质属于变量，在赋值、函数形参时为副本，结果两个数组操作互不干涉。

	kk,ff:=toolNya2.Exp2(a);
	fmt.Printf("kk=%d ff=%d a=%d\n",kk,ff,a);
	//gg Exp3内 y;hh Exp3内 x;b 原切片
	b := [5] int{0,9,8,7,6} //这是数组
	gg,hh:=toolNya2.Exp3(b);
	fmt.Printf("gg=%d hh=%d b=%d\n",gg,hh,b);
	//gg Exp3内 y;hh Exp3内 x;b 原数组

}
