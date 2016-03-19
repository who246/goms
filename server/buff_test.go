package server
import (
    "fmt" 
	"testing" 
	"bytes"
)
func Test_Buff(t *testing.T) {
//	s := []byte(" world")
 	s := make([]byte,2,2) 
	s[1] = 1
    buf := bytes.NewBuffer(s)
    fmt.Println(buf.String())  //buf.String()方法是吧buf里的内容转成string，以便于打印
    buf.Write([]byte("world")) //将s这个slice写到buf的尾部
    fmt.Println(buf.String())  //打印 hello world
}