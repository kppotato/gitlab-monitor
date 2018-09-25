package biz

import (
	api "github.com/kppotato/gitlabMonitorUI/dao"
	"github.com/kppotato/gitlabMonitorUI/model"
)
func GetProjects() []*model.WrapeProject{
	projects,err := api.ListProjects()
	if err !=nil{
		return nil
	}
	result :=make([]*model.WrapeProject,len(projects))
	for _,x:=range projects{
		project:=model.NewWrapeProject(x)
		result=append(result, project)
		piplinemap,err :=api.GetPiplines(x.ID)
		if err !=nil{
			continue
		}
		if piplinemap == nil{
			continue
		}
		for _,v :=range piplinemap{
			line:= model.NewWrapePipline(v)
			project.Piplines = append(project.Piplines,line)
			jobs,err :=api.GetJobs(x.ID,v.ID)
			if err !=nil{
				continue
			}
			line.Jobs=jobs
		}
	}
	//log.Info("len result:",len(result),cap(result))
	var a=make([]int,0)
	for index,x:=range result{
		if x == nil{
			continue
		}
		if x.Piplines== nil || len(x.Piplines)==0{
			continue
		}
		a=append(a,index)
	}
	newresult:=make([]*model.WrapeProject,len(result)-len(a))
	for _,i:=range a{
		newresult=append(newresult,result[i])
	}
	return newresult
}
