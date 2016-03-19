package socket

import(
 "log"
 "time"
"fmt"
)
 
type receivPack struct {
	group string
	topic string
	msg string
}
type sendPack struct{
	msgType int
	msg string
}

type executer interface {
	readPacket()
    writePacket(pack *sendPack)
}

type sender struct {
	 h Host
}
func (s *sender) readPacket(){
	 
	 t,err := s.h.topic()
	 if err != nil{
		log.Fatal(err)
		return
	 }
	 g,err := s.h.group()
	 if err != nil {
		log.Fatal(err)
		return
	 }
	 b,err := s.h.body()
	 if err != nil {
		log.Fatal(err)
		return
	 }
	 p := &receivPack{group:g,topic:t,msg:string(b)}
	 pa <- p
	
}
func (s *sender) writePacket(pack *sendPack){
	 
}

type resceive struct {
	 h Host
}

func (s *resceive) writePacket(pack *sendPack){
    //s.h.conn.Write(msg)
	fmt.Println(pack.msg)
}
func (s *resceive) readPacket(){
	for{
	   op,err := s.h.op()
	   if err != nil {
		log.Fatal(err)
	   }
	  switch op {
		case PONGMESSGE:
		   s.h.setReadDeadline(time.Now().Add(TIMEOUT))
	  }
	   
	}
}
 