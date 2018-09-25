package model

import (
	"github.com/kppotato/gitlabMonitorUI/go-gitlab"
)


type WrapeProjectSort []*WrapeProject

func (this WrapeProjectSort) Len() int {   // 重写 Len() 方法
	return len(this)
}
func (this WrapeProjectSort) Swap(i, j int){  // 重写 Swap() 方法
	this[i], this[j] = this[j], this[i]
}
func (this WrapeProjectSort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	//last_activity_at
	if this[j] == nil{
		return false
	}
	if this[i] == nil{
		return false
	}
	return this[j].LastActivityAt.Before(this[i].LastActivityAt)
}
type WrapeProject struct {
	*gitlab.Project
	Order						int
	Piplines					[]*WrapePipline
}

func NewWrapeProject(project *gitlab.Project) *WrapeProject{
	return &WrapeProject{Order:0,Project:project,Piplines:make([]*WrapePipline,0)}
}


type WrapePipline struct{
	*gitlab.Pipeline
	Jobs			[]*gitlab.Job
}

func NewWrapePipline(pipline *gitlab.Pipeline) *WrapePipline{
	return &WrapePipline{Pipeline:pipline}
}