// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/kppotato/gitlabMonitorUI/g"
	"github.com/kppotato/gitlabMonitorUI/biz"
	_ "github.com/kppotato/gitlabMonitorUI/cron"
	log "github.com/Sirupsen/logrus"
	"github.com/kppotato/gitlabMonitorUI/cron"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gitlabMonitorUI",
	Short: "gitlab-ci monitor web",
	Long: `gitlab-ci monitor web`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		g.GITLAB_CI_API_URL = viper.GetString("apiurl")
		g.GITLAB_CI_API_TOKEN = viper.GetString("apitoken")
		log.Info("gitlab api url:",viper.GetString("apiurl"))
		log.Info("gitlab api token:",viper.GetString("apitoken"))
		log.Info("gitlab-monitor start http server:0.0.0.0:",viper.GetString("port"))
		if g.GITLAB_CI_API_TOKEN =="" || g.GITLAB_CI_API_URL==""{
			log.Error("apiurl or apitoken not empty please -h for help")
			os.Exit(1)
		}
		g.Oninit()
		cron.Oninit()
		biz.StartHttp()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	flagset := RootCmd.Flags()
	flagset.StringVar(&g.GITLAB_CI_API_URL,"apiurl","","gitlab-ci api url")
	flagset.StringVar(&g.GITLAB_CI_API_TOKEN,"apitoken","","gitlab-ci api token")
	flagset.IntVar(&g.Port,"port",8080,"gitmonitor port")
	viper.BindPFlags(flagset)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.BindEnv("apitoken")
	viper.BindEnv("apiurl")
	viper.BindEnv("port")
}
