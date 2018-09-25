package biz

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	//log "github.com/Sirupsen/logrus"
	"github.com/kppotato/gitlabMonitorUI/controller"
	"github.com/kppotato/gitlabMonitorUI/g"
	"strconv"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin/json"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kppotato/gitlabMonitorUI/data"
	"strings"
)

func StartHttp()  {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Delims("[[","]]")
	//r.LoadHTMLGlob("./template/*")

	t, _ := loadTemplate()
	r.SetHTMLTemplate(t)
	//r.Static("/static", "./static")
	fs:=assetfs.AssetFS{Asset:data.Asset,AssetDir:data.AssetDir,AssetInfo:data.AssetInfo,Prefix:"static"}
	r.StaticFS("/static",http.FileSystem(&fs))

	r.GET("", func(c *gin.Context) {
		result:=GetProjects()
		b,err := json.Marshal(result)
		if err !=nil{
			log.Error("GetProjects error:",err)
			c.HTML(http.StatusOK, "index.html",gin.H{"jsonobj":"[]"})
			return
		}
		c.HTML(http.StatusOK, "index.html",gin.H{"jsonobj":string(b)})



	})
	r.GET("/ws",controller.WsHandler())
	r.Run("0.0.0.0:"+strconv.Itoa(g.Port))
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	names,_:=data.AssetDir("template")
	for _,name := range names{
		b,err :=data.Asset("template/"+name)
		if err !=nil{
			log.Info(err)
			continue
		}

		if  !strings.HasSuffix(name, ".tmpl") && !strings.HasSuffix(name, ".html"){
			continue
		}

		t, err = t.New(name).Delims("[[","]]").Parse(string(b))
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	return t, nil
}