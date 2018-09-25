package cron

import (
	"time"
	"github.com/kppotato/gitlabMonitorUI/g"
	"github.com/kppotato/gitlabMonitorUI/controller"
	log "github.com/Sirupsen/logrus"
	"github.com/kppotato/gitlabMonitorUI/biz"
	"encoding/json"
)

func Oninit(){
	go GetGitlab()
	go controller.HandlerWesocket()
}

func GetGitlab()  {
	for{
		select {
			case <-g.Exit:
				break
			default:
				log.Info("cron getlibtal")
				result :=biz.GetProjects()
				//sort.Slice(result, func(i, j int) bool {
				//	if result[i] == nil{
				//		return false
				//	}
				//	if result[j] == nil{
				//		return true
				//	}
				//	return result[j].LastActivityAt.Before(result[i].LastActivityAt)
				//})

				b,err :=json.Marshal(result)
				if err != nil{
					continue
				}
				g.Data <- string(b)
		}
		<-time.After(time.Duration(g.Interval)*time.Second)
	}
}