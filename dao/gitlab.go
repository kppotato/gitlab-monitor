package biz

import (
	"github.com/kppotato/gitlabMonitorUI/go-gitlab"
	"github.com/kppotato/gitlabMonitorUI/g"
	log "github.com/Sirupsen/logrus"
	"time"
)

func ListProjects() ([]*gitlab.Project,error){
	opt :=&gitlab.ListProjectsOptions{
		OrderBy:gitlab.String("last_activity_at"),
	}
	projects,_,err :=g.GitClient.Projects.ListProjects(opt,nil)
	if err != nil{
		log.Errorf("get project error: %s",err)
		return nil,err
	}
	return projects,err
}
func GetPiplines(projectid int)(map[string]*gitlab.Pipeline,error){
	piplinesOpt := &gitlab.ListProjectPipelinesOptions{
		ListOptions:gitlab.ListOptions{Page:0,PerPage:1},
		//OrderBy:gitlab.String("id"),
		//Sort:gitlab.String("desc"),
	}
	branchopt := &gitlab.ListBranchesOptions{}
	branchs,_,err := g.GitClient.Branches.ListBranches(projectid,branchopt)
	if err != nil{
		log.WithField("projectId",projectid) .Errorf("get ListBranches error: %s",err)
		return nil,err
	}
	//g.GitClient.Tags.ListTags(projectid,)
	result :=make(map[string]*gitlab.Pipeline)
	for _,x :=range branchs{
		piplinesOpt.Ref = &x.Name
		list,_,err := g.GitClient.Pipelines.ListProjectPipelines(projectid,piplinesOpt)
		if err !=nil{
			log.WithField("projectId",projectid) .Errorf("get piplineslist error: %s",err)
			continue
		}
		if len(list)>0{
			for _,p:=range list{
				pipline,_,err :=g.GitClient.Pipelines.GetPipeline(projectid,p.ID)
				if err !=nil{
					log.WithField("projectId",projectid).WithField("piplineid",p.ID) .Errorf("get piplineslist error: %s",err)
					continue
				}
				if value,ok:=result[x.Name];ok{
					if value.UpdatedAt.Before(*pipline.UpdatedAt){
						result[x.Name] = pipline
					}
				}else{
					result[x.Name] = pipline
				}
			}
		}
	}
	d,_:=time.ParseDuration("-24h")
	for k,v :=range result{
		if v.FinishedAt==nil{
			continue
		}
		if v.FinishedAt.Before(time.Now().Add(d)){
			delete(result,k)
		}
	}
	return result,nil
}
func GetJobs(projectid,piplineid int)([]*gitlab.Job,error){
	jobs,_,err := g.GitClient.Jobs.ListPipelineJobs(projectid,piplineid,nil)
	if err != nil{
		log.WithField("projectId",projectid).WithField("piplineid",piplineid) .Errorf("get jobs error: %s",err)
		return nil,err
	}
	return jobs,nil
}
