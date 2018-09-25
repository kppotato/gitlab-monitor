package util

import (
	"sync"
	"github.com/gorilla/websocket"
	"fmt"
)

type Hub struct {
	sync.RWMutex
	Data     map[*websocket.Conn]struct{}
}

func Newhub() *Hub{
	return &Hub{Data:make(map[*websocket.Conn]struct{},100)}
}

func (this *Hub) Add(c *websocket.Conn){
	this.Lock()
	defer this.Unlock()
	fmt.Print("add conn",c.RemoteAddr())
	this.Data[c]= struct{}{}
}
func (this *Hub) Remove(c *websocket.Conn)  {
	this.Lock()
	defer  this.Unlock()
	if _,ok :=this.Data[c];ok{
		fmt.Println("remove conn",c.RemoteAddr())
		delete(this.Data,c)
	}
}
func (this *Hub) GetALL() map[*websocket.Conn]struct{}{
	this.RLock()
	defer this.RUnlock()
	return this.Data
}


