package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Arith int

type Args struct
{
	A ,B int
}

type AnsList struct {
	Ans  int
	Ans2 string
}
type StateList struct {
	numA ,numB float64
	Methom string
	Authors []string
}


func (this *Arith) Multiply (args *Args , reply *AnsList) error {

	(*reply).Ans = args.A * args.B;
	s:=&StateList{
		float64(args.A),float64(args.B),
		"Multiply",
		[]string{ "plague", "cat"},
	}
    b,err :=json.Marshal(s);
    if err != nil{fmt.Println("ERR Multiply 01"); return nil;}
	(*reply).Ans2 =string(b);

	return nil;
}

func RunSever(IPaddress string) {

	arith:=new(Arith);
	rpc.Register(arith);

	tcpAddr, _ := net.ResolveTCPAddr("tcp", IPaddress)
	l, _ := net.ListenTCP("tcp", tcpAddr)

	fmt.Println("Start Listen")
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)

	}
}


func main() {

	go RunSever("127.0.0.1:9000");

	client, err := rpc.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("error: ", err)
	}

	args:=&Args{7,8}

	reply:=new(AnsList)
	err = client.Call("Arith.Multiply", args, reply)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Arith  %d*%d=%d\n",args.A,args.B ,reply.Ans )
	fmt.Printf("JSONNote:%s \n",reply.Ans2)
	b:=[]byte(reply.Ans2)
	s:=new(StateList)
	err = json.Unmarshal(b,&s)
	if err !=nil{log.Println(err)}
	fmt.Println("<==== Decode ====>\n");
	fmt.Printf("Methom: %s\n  \n",s.Methom )
	fmt.Printf("Authors: \n" )
	//sa 就是个int 0---->len
	for sa :=range s.Authors {
		fmt.Printf(" %d : %s\n", sa , s.Authors[sa]) ;
	}
	//未知情况下 转换json
	fmt.Println("<====Unkown Json Decode ====>\n");
	var r interface{} //用空接口接住数据
	err = json.Unmarshal(b,&r);if err !=nil{log.Println(err)}
	//此时 r对象 内部应该接受为map[string] interface
     statemap ,ok := r.(map[string]interface{})
     if ok{
		 fmt.Println("Find is Map:\n");
     	author,ok2 :=statemap["Authors"].([]interface{});

     	if ok2{
			fmt.Println("Find Map[\"Authors\"] is Array is []interface\n");
			_,ok3 :=author[0].(string)
			ok4:=true;
			switch author[0].(type) {
				case string:
					ok4=true;
				default:
					ok4=false;
				}
			author,ok5 :=statemap["Authors"].([]string);
			fmt.Printf("But Map[\"Authors\"] is not []string, ok5 is %g\n ",  ok5);
			if ok3&&ok4 {
				fmt.Println("Find Map[\"Authors\"][num] is string\n");
			    for sa2 ,sa2v := range author{
					fmt.Printf(" %d : %s\n", sa2 , sa2v);
				}
			}
		}
     }
}




