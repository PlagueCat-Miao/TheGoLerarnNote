package main

import "fmt"



type Rect struct {
	x , y float64
	Width ,Height float64
	ID int //这个会覆盖默认Rect.ID从Rect.Base.ID到Rect.ID
	Base

}
type Base struct{
	ID int
}


func (this *Rect ) Area() float64 {
	(*this).Height=9;  //此处 (*p).a 与 p.a 发生同义了 https://bbs.csdn.net/topics/390974889?page=1
	this.Width=8;
	return this.Height*this.Width
}
func (this Rect ) Area2() float64 {
	//(*this).Height=9;  //此处 (*p) 非法
	this.Width=8;
	return this.Height*this.Width
}
/*func NewRect() *Rect{
	return &Rect{-44,-13,6,6,}
}*/
/*
func main() {
    rect1 := &Rect{1,2,3,4,1,Base{4}};

	fmt.Printf("%f %f \n", rect1.Area(),rect1);

	rect3 :=new(Rect);//都是零
	rect3.ID=2;
	rect3.Base.ID=3;
	fmt.Printf("ok %f %f %f \n", rect1,rect3,rect3.Area2());//rect3.Height,rect3.Width);



}
*/

type inter int
type LessAddoneer interface {
	Less() bool
	Addone()
}
type LADer interface {
	LessAddoneer
	de()
}
func (b *inter) Addone(){
	*b +=  1;
}
func (b inter) Less() bool {
	return b < 4;
}
type miao struct
{
	nya int
}

func (b *miao) de()  {
	(*b).nya--;
}
func (b *miao) Addone(){
	b.nya +=  1;
}
func (b miao) Less() bool {
	return b.nya < 4;
}

type IOstream interface {
	Input()
	Output()
}
func main() {
	var a inter=2;
	var b LessAddoneer = &a;
	b.Addone()
	fmt.Printf("%d %c\n" ,a,b.Less());
	var c LADer =new(miao);

   if _, ok:= c.(LessAddoneer) ; ok {
       fmt.Printf("ok  c 满足LessAddoneer接口\n");
   }
	if _, ok:= b.(LADer) ; ok {

	} else{
		fmt.Printf("no  b 不满足LADer接口\n");
	}
	if _, ok:= c.(*miao) ; ok {
		fmt.Printf("ok  c 满足LessAddoneer接口\n");
	}
}


