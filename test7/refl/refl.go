package main

import (
	"fmt"
	"reflect"
)
type Miao struct {
	NumId int
	LittleMiao
}
type Miao2 struct {
	NumId2 int
	LittleMiao
}
type LittleMiao struct {
	Nya string
}
func (this *Miao) say() string {
	return this.Nya;
}
func (this *Miao2) say() string {
	return this.Nya;
}
type ani interface {
	say() string
}
func main() {
	var x float64
	p:=reflect.ValueOf(&x)
	x= 4.3;
	fmt.Println("typy of p:",p.Type())
	fmt.Println("p settability",p.CanSet())//指针不能被set

	v:=p.Elem()//得到指针下的对象
	fmt.Println("v settability",v.CanSet())
	fmt.Println("Interface returns v's current value as an interface{}. get:",v.Interface())

	z:=reflect.ValueOf(x)
	fmt.Println("typy of z:",z.Type())
	fmt.Println("z settability",z.CanSet())//值的副本 不能被set


	//<========test========>
	fmt.Println("p : v : x : z")
	fmt.Println(p,":",v,":",x,":",z)
	fmt.Println("v set 7.1")
	v.SetFloat(7.1)
	fmt.Println(p,":",v,":",x,":",z)

	fmt.Println("p 是 x的指针的value ; 指针不能被set")
	fmt.Println("v 是 x的指针value(p)所指向的对象的value ; \n   与x等价 可set（但其是value,不能像正常值、同名一样操作，如用=赋值）")
	fmt.Println("z 是 x的副本的value ; 副本不能被set 且独立于x 不受x的变化影响")

	miao:=Miao{4,LittleMiao{"mewo"}}
	m:=reflect.ValueOf(&miao).Elem()//Elem()解地址了
	typeOfM:=m.Type()//成员的名（type）
	for i:=0; i<m.NumField();i++{
		f:=m.Field(i)//储存成员的值
		fmt.Printf("%d : %s %s = %v\n",i,
			typeOfM.Field(i).Name , f.Type() , f.Interface())
	}
	fmt.Println("<===Unknow Data========>")
	var RR interface{}
	RR = miao;
	R:=reflect.ValueOf(&RR).Elem()//这里已经解地址了

	typeOfR:=R.Type()//成员的名（type）
	fmt.Println("显然如果是空指针 type只能识别为空指针",typeOfR);
	_ ,ok := RR.(map[string]interface{})
	if ok{
		fmt.Println("Find is Map:\n");
	}else if _ ,ok2 := RR.(Miao2); ok2{
		fmt.Println("Find is Miao2:\n");

	}else if mm ,ok3 := RR.(Miao); ok3{
		fmt.Println("Find is Miao:");
		fmt.Println("但是能通过RR.(Miao)来得到正确的类 并赋值");
		fmt.Println("mm.Nya is" ,mm.Nya);

	}
	if mm ,ok4 := RR.(ani); ok4{
		fmt.Println("mm.Nya is" ,mm.say());
	} else{
		fmt.Println("不能通过RR.(ani)来查看是否满足某个接口");
	}

	/*
	for i:=0; i<2;i++{
//		f:=R.Field(i)//储存成员的值
		fmt.Printf("%d : %s %s = %v\n",i,
			typeOfR.Field(i).Name , f.Type() , f.Interface())
	}
    */

}
