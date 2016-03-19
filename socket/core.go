package socket

import(
//	 "fmt"
	"sync"
	 "errors"
	"log"
)

var receivers map[string]map[string][]executer
var pa chan *receivPack
var mutex sync.Mutex
func init(){
	receivers = make(map[string]map[string][]executer)
	pa = make(chan *receivPack)  
}

func getReceivers(topic string)  (map[string][]executer,error){
	g,ok := receivers[topic]
	if ok {
		return g,nil
	}else{
		return nil,errors.New("sender not exit")
	}
	
}
func register(ex executer,topic,group string){	
      mutex.Lock()
      defer mutex.Unlock()
	  g, ok := receivers[topic] 
	  if ok {
		 es,ok:=g[group]
		 if ok {
			es = append(es,ex)
		 }else{
		 g[group] = es
		 receivers[topic] = g
		 }
	  } else {
		es := []executer{ex}
	 	g := make(map[string][]executer)
		g[group] = es
		receivers[topic] = g
	  }
}

func WriteTask(){
	 for {
		select {
			case p, ok := <-pa :
			if ok {
			g,err := getReceivers(p.topic)
			if err != nil {
				log.Fatal("receiv error",err)
				break
			}
			 for _, v := range g {
			   go send(v,p)
			 }
			}
		}
	 }
}
func send(exs []executer,p *receivPack){
	  ex := exs[0]
	 //ex,err := getReceivers(p.topic,p.group)
	//not find
//	 if err != nil{
//		log.Fatal("send error",err)
//		return
//	 }
	  ex.writePacket(&sendPack{msgType:TEXTMESSGE,msg:p.msg})
}