package cmd

import (
	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var wechatCmd = &cobra.Command{
	Use:   "wechat",
	Short: "Post Jira filtered result to Wechat",
	Run:   wechat,
}

func wechat(cmd *cobra.Command, args []string) {
	readConfig()
	jiraFilter()
}

func jiraFilter() {
	tp := jira.BasicAuthTransport{
		Username: jiraConf.Username,
		Password: jiraConf.Password,
	}
	jClient, err := jira.NewClient(tp.Client(), jiraConf.Host)
	if err != nil {
		log.Fatal("Fail to connect to JIRA server, the error is: %v", err)
	}

	for _, filter := range jiraConf.Filter {
		log.Info(filter.Name)
		opt := &jira.SearchOptions{}
		issues, _, err := jClient.Issue.Search(filter.Jql, opt)
		if err != nil {
			log.Fatalln(err)
		}
		for _, issue := range issues {
			log.Info(jiraConf.Users[issue.Fields.Assignee.Name], issue.Fields.Summary)
		}
	}

}
