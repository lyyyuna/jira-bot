package internal

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
	"net/url"
)

func RunFilter(jiraConf TomlConfig) []*SectionStats {
	tp := jira.BasicAuthTransport{
		Username: jiraConf.Username,
		Password: jiraConf.Password,
	}
	jClient, err := jira.NewClient(tp.Client(), jiraConf.Host)
	if err != nil {
		log.Fatal("Fail to connect to JIRA server, the error is: %v", err)
	}

	jiraStats := make([]*SectionStats, 0)
	for _, filter := range jiraConf.Filter {
		section := &SectionStats{
			Name:  filter.Name,
			Jql:   filter.Jql,
			Url:   jiraConf.Host + "issues/?jql=" + url.QueryEscape(filter.Jql),
			Users: make(map[string]*UserStats, 0),
			Cnt:   0,
			Split: filter.Split,
		}
		jiraStats = append(jiraStats, section)
		log.Infof("Filter name: %v", filter.Name)
		opt := &jira.SearchOptions{}
		issues, _, err := jClient.Issue.Search(filter.Jql, opt)
		if err != nil {
			log.Fatalln(err)
		}
		section.Cnt = len(issues)
		// Skip if split is false
		// true will display the results grouped by name
		if filter.Split == false {
			continue
		}

		for _, issue := range issues {
			if issue.Fields.Assignee == nil {
				log.Errorf("This issue has no assignee, the issue is: %v", issue.Fields.Summary)
				continue
			}
			// log.Info(issue.Fields.Assignee.Name, issue.Fields.Summary)
			jiraName := issue.Fields.Assignee.Name
			if _, ok := section.Users[jiraName]; ok {
				section.Users[jiraName].Cnt++
			} else {
				jql := fmt.Sprintf("Assignee = %v and %v", jiraName, section.Jql)
				section.Users[jiraName] = &UserStats{
					Jql: jql,
					Cnt: 1,
					Url: jiraConf.Host + "issues/?jql=" + url.QueryEscape(jql),
				}
			}
		}
	}

	return jiraStats
}
