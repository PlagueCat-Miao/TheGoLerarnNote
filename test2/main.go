package main

import "fmt"

func main() {
	var (
		v1 int
		v2 string
	)
	const (
		AA = iota;
		BB
		CC
		DD
		EE
		end
	)
	v1=3;
	v2="uihih";
	fmt.Println("%s %d",v2,v1," ",AA,  end);
    if (1==2 || 3!=4){
    	fmt.Println(" 1==2||3!=4 =>",1==2 ,"||", 3!=4)
	} else
    {
		fmt.Println(" NO");
	}
	ch := v2[4];
	ch =v2[0];
	var ch3 byte;
	ch3 =ch;
	fmt.Printf("%s the first :%c len is %d" ,v2,ch3, len(v2));
 for i:=0 ;i<len(v2);i++{
	 fmt.Printf("%c\n" ,v2[i]);
 }

    slice :=make ([]int ,3,5);
	slice=append(slice, 10,11,12,13,14,15,16,17);
	fmt.Printf("slice len3,cap 5=>len=%d,cap=%d \n",len(slice),cap(slice));
    var mmap = make( map[string] int)
    mmap ["miao"]= 4;
    kk, _ :=mmap["miao"];
    fmt.Printf("mmap %d-miao",kk);
	/**/
}
