package main
import (
	"fmt"
	"net"
	"os"
	"io"
	"bytes"
)

func tcp_main() {
	if len(os.Args) !=2{
		fmt.Println("Args error")
		os.Exit(0);
	}
	serverIp :=os.Args[1]
	conn, err :=net.Dial("tcp",serverIp)
	checkerror(err);
	conn.Write([]byte("miao\000") ) //为了和C++对接 永远不要忘记 \0 要不c++的recv会读出len长个乱码

	result , err := readFully(conn);
	checkerror(err);
	fmt.Println(string(result),result[0:9] ); //strin() 遇到/000会停止喵
	os.Exit(0);
}

func checkerror(err error) {
	if err != nil {
	fmt.Fprintf(os.Stderr,"Err : %s", err.Error() );
		os.Exit(1)
    }
}
func readFully(conn net.Conn)([]byte ,error){
    defer conn.Close()
	 result :=bytes.NewBuffer(nil);
	 var buf [512]byte
	 miaoread:
	 for{
		 n,err:=conn.Read(buf[0:]) //在C++中 没有io.EOF 一般用/000 /0结尾
		 fmt.Println(n," ",buf[0:n] );
		 result.Write(buf[0:n])  //你遇到\0还会停啊？

		 if err != nil{
		 	if err == io.EOF {  //C++的send 似乎发不出io.EOF
		 		break miaoread;
			}

			 fmt.Println("bad not detect EOF now err!=nill"); //判定到别的错误了
			 return result.Bytes(), err
		 } else{
			 fmt.Println("bad not detect EOF now err==nill"); //没有判定到结束符
			 return result.Bytes(), err
		 }

	 }
	 	return result.Bytes() ,nil
}