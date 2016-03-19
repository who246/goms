package socket

import(
	 "fmt"
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

func getReceivers(topic,group string)  (executer,error){
	g,ok := receivers[topic]
	if ok {
		t,ok := g[group]
		if ok {
			return t[0],nil
		}else{
			return nil,errors.New("sender not exit")
		}
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
			go send(p)
			}
		}
	 }
}
func send(p *receivPack){
	 fmt.Println(p.msg)
	 _,err := getReceivers(p.topic,p.group)
	//not find
	 if err != nil{
		log.Fatal("not find receiver")
		return
	 }
//	  ex.write([]byte(p.msg))
}