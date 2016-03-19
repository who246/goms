package main

import (
//	"fmt"
    "log"
     "net" 
	"flag"
	"strconv"
//	 "io"
	"github.com/who246/goms/socket"
)
 

//func read(conn net.Conn) string{
//	 BufLength := 10
//     buffer  := make([]byte, BufLength)
//     data := make([]byte, 0)
//    // m := 0
//	for {
//		n, err := conn.Read(buffer)
//	//	m = n+m;
//		if err != nil && err != io.EOF {
//			fmt.Println(err)
//				break
//			}
		
//		data = append(data, buffer[:n]...)
//			if n != BufLength {
//				break
//			}
//		}
//		fmt.Println(len(data))
 
//	return string(data)
//}
//func readOp(conn net.Conn) int{
//	data := make([]byte, 1,1)
//	 conn.Read(data)
//	return int(data[0])
//}

//func handleConnection(conn net.Conn) {
   
	
//	 for {
//		 fmt.Println(readOp(conn))
//		 fmt.Println(read(conn))
//	 }

//		//fmt.Println("data:", string(data)) 
//		//fmt.Fprintf(conn, "hello client\n")
//		conn.Close()
//}

func main(){
	port := *flag.Int("p",8808,"Please enter the correct port")
	flag.Parse()
	ports := strconv.Itoa(port)
	l, err := net.Listen("tcp", ":"+ports)
    if err != nil {
        panic(err)
    }
	go socket.WriteTask()
	for{
	   con,err := l.Accept()
	   if err != nil{
		log.Fatal("get connection error",err)
		continue
	   }
	   s := socket.NewServer(con)
	   go s.HandleConnection()
	}
}