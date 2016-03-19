package socket

import(
 "log"
 "time"
//"fmt"
"bytes"
)
const(
	pingPeriod = (30 * time.Second) 
)
type receivPack struct {
	topic string
	msg []byte
}
type sendPack struct{
	msgType byte
	msg []byte
}

type executer interface {
	readPacket()
    writePacket(pack *sendPack) error
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

	 b,err := s.h.body()
	 if err != nil {
		log.Fatal(err)
		return
	 }
	 p := &receivPack{topic:t,msg:b}
	 pa <- p
	
}
func (s *sender) writePacket(pack *sendPack) error {
 return nil	 
}

type resceive struct {
	 h Host
}

//发送数据
func (s *resceive) writePacket(pack *sendPack) error {
	b := make([]byte,1,1) 
 	b[0] =  pack.msgType
	buf := bytes.NewBuffer(b)
 	_,err :=  buf.Write(pack.msg)
	if err != nil {
		return err
	}
	_,err = s.h.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
//接收接收者信息
func (s *resceive) readPacket(){
	for{
	   op,err := s.h.op()
	   if err != nil {
		log.Fatal(err)
		return
	   }
	  switch op {
		case PONGMESSGE:
		   s.h.setReadDeadline(time.Now().Add(TIMEOUT))
	  }
	   
	}
}
//心跳
func (s *resceive) heartbeat(){
	ticker := time.NewTicker(pingPeriod)
	defer func() {		
		ticker.Stop()		
	}()
	for {
		select {
			case <-ticker.C: 
			 if err := s.writePacket(&sendPack{msgType:PINGMESSGE}); err != nil {
				log.Fatal("PINGMESSGE error",err)
				return
			 }
		}
	}
}