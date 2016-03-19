package server
import (
    "fmt"
    "net"
  //  "bufio"
	"testing"
 //	"time"
	"bytes"
)

func Test_Consumer(t *testing.T){
	
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