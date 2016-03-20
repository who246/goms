package server
import (
    "fmt"
    "net"
  //  "bufio"
	"testing"
 //	"time"
	"bytes"
	"github.com/who246/goms/socket"
)

func read(c net.Conn) ([]byte,error) {
	b := make([]byte,1,1) 
	_,err := c.Read(b)
	if err!=nil{
		return nil,err
	}
	return b,nil
}

func Test_Consumer(t *testing.T){
	conn, err := net.Dial("tcp", ":8808")
	if err != nil {
        panic(err)
    }
 	b := make([]byte,1,1) 
 	b[0] = 1
	buf := bytes.NewBuffer(b)
	conn.Write(buf.Bytes())
	for{
		b,err := read(conn)
		if err != nil{
			fmt.Println(err)
			continue
		}
		if b[0] == socket.PINGMESSGE {
			pong :=  []byte{byte(socket.PONGMESSGE)}
			conn.Write(pong)
		}
	}
      conn.Close()
}
func Test_Producer(t *testing.T) {
    conn, err := net.Dial("tcp", ":8808")
    if err != nil {
        panic(err)
    }
 	b := make([]byte,1,1) 
 	b[0] = 1
	words := `hello consumer`
	buf := bytes.NewBuffer(b)
 	buf.Write([]byte(words))
	fmt.Println(buf.String())
	conn.Write(buf.Bytes())
      conn.Close()
}