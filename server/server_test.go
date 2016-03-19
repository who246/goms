package server
import (
    "fmt"
    "net"
  //  "bufio"
	"testing"
 //	"time"
	"bytes"
)
func Test_Server(t *testing.T) {
    conn, err := net.Dial("tcp", ":8808")
    if err != nil {
        panic(err)
    }
 	b := make([]byte,1,1) 
 	b[0] = 1
	words := `2ddddddd\nsdsd`
	buf := bytes.NewBuffer(b)
 	buf.Write([]byte(words))
	fmt.Println(buf.String())
//   for{
	conn.Write(buf.Bytes())
//	time.Sleep(10 * time.Second)
//	}
	//words = "2hello server2"
	//conn.Write([]byte(words))
	//conn.Close()
  //  data, err := bufio.NewReader(conn).ReadString('\n')
//    if err != nil {
//        panic(err)
//    }
 //   fmt.Printf("%#v\n", data)
      conn.Close()
}