package socket

import(
	"net" 
	"time"
	"log"
	"errors"
)
const(
	TIMEOUT = 2*60 * time.Second
)



func NewServer(con net.Conn) *Server{
	h := Host{conn:con}
	h.setReadDeadline(time.Now().Add(TIMEOUT))
	return &Server{h:h}
}
type Server struct{
	 h Host
	 
}

func (s *Server)HandleConnection() {
	  defer func() {
	   s.h.close()
	  }() 
	  id ,err := s.h.identity()
	  if err != nil {
		log.Fatal(err)
		return 
	  }
	
	ex,err := s.getStrategy(id)
	if err != nil {
		log.Fatal(err)
		return 
	  }
	
	ex.readPacket()
	
}



func (s *Server) getStrategy( id int) (ex executer,e error) {
	if id == SENDER {
		ex = &sender{h:s.h}
		e = nil 
	}else if id == RESCEIVE {
		r := &resceive{h:s.h}
		ex = r
		t,err := s.h.topic()
		if err != nil{
			return nil,err
		}
		g,err := s.h.group()
		if err != nil{
			return nil,err
		}
		//注册
		register(ex,t,g)
		//启动心跳
		go r.heartbeat()
		e = nil 
	}else{
		e = errors.New("not correct packet")
	}
	return
}