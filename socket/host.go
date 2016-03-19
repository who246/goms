package socket

import(
	"net" 
	"time"
	"io"
)

type Host struct{
	conn net.Conn
}
func (h * Host)  setReadDeadline( t time.Time){
	 h.conn.SetReadDeadline(t)
}
func (h * Host) close(){
	 h.conn.Close()
}
func (h * Host) topic() (string,error){
	 return "topic",nil
}
func (h * Host) group() (string,error){
	 return "group",nil
}
func (h * Host) identity() (int,error){
	buf := make([]byte, 1)
	if _, err := io.ReadFull(h.conn, buf); err != nil {
        return 0,err
    }
	return int(buf[0]),nil
}
func (h * Host) op() (int,error){
     buf := make([]byte, 1)
	if _, err := io.ReadFull(h.conn, buf); err != nil {
        return 0,err
    }
	return int(buf[0]),nil
}
func (h * Host) body()([]byte,error){
    return read(h.conn)
}

func  read(conn net.Conn) ([]byte,error){
	 BufLength := 1024
     buffer  := make([]byte, BufLength)
     data := make([]byte, 0)
	for {
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
				return nil,err
			}
		
		data = append(data, buffer[:n]...)
			if n != BufLength {
				break
			}
		}
	  return data,nil
}