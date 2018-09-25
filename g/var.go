package g

import (
	"github.com/kppotato/gitlabMonitorUI/go-gitlab"
)

var(
	GITLAB_CI_API_URL				string
	GITLAB_CI_API_TOKEN				string
	GitClient						*gitlab.Client
	Exit							chan struct{}
	Data							chan string
	Interval						int =10
	Port							int
)

func init(){
	Exit=make(chan struct{},1)
	Data = make(chan string,1)
	Port = 8080
}
func Oninit()  {
	GitClient =gitlab.NewClient(nil,GITLAB_CI_API_TOKEN)
	GitClient.SetBaseURL(GITLAB_CI_API_URL)
}