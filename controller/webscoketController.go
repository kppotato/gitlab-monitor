package controller

import (
	"github.com/gorilla/websocket"
	"github.com/kppotato/gitlabMonitorUI/util"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kppotato/gitlabMonitorUI/g"
	log "github.com/Sirupsen/logrus"
)

var(
	WebsocketClient =util.Newhub()
	Wsupgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		// 取消ws跨域校验
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
func WsHandler() func(context *gin.Context){
	return func(c *gin.Context) {
		var conn *websocket.Conn
		var err error
		conn, err = Wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		WebsocketClient.Add(conn)
	}
}
func HandlerWesocket(){
	for{
		select{
		case data:= <- g.Data:
			log.Info("send data:",len(data))
			conns := WebsocketClient.GetALL()
			log.Info("conns:",len(conns))
			for conn,_ := range conns{
				err := conn.WriteMessage(websocket.TextMessage,[]byte(data))
				if err != nil{
					WebsocketClient.Remove(conn)
					continue
				}
				log.Info("websocket:",conn.RemoteAddr())
			}
		}
	}
}