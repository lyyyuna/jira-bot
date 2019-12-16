package cmd

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	. "jira-bot/internal"
	"text/template"
)

const jiraTemplate = `
# {{.Info.Project}}
{{range .Sections}}
## {{.Name}}
总数： [{{.Cnt}}]({{.Url}})
{{if .Split}}
{{range $k, $v := .Users}}{{$k}} 有 [{{$v.Cnt}}]({{$v.Url}})
{{end}}
{{end}}

{{end}}
`

var wechatCmd = &cobra.Command{
	Use:   "wechat",
	Short: "Post Jira filtered result to Wechat",
	Run:   wechat,
}

func wechat(cmd *cobra.Command, args []string) {
	readConfig()
	sections := RunFilter(jiraConf)
	content := markdown(sections)

	c := NewWechatClient(jiraConf.Wxkey)
	c.SendToWechat(content)
}

func markdown(sections []*SectionStats) string {
	tmpl, err := template.New("jira").Parse(jiraTemplate)
	if err != nil {
		log.Fatalf("Error while parsing template, the error is: %v", err)
	}

	var buf bytes.Buffer
	items := struct {
		Info     TomlConfig
		Sections []*SectionStats
	}{
		Info:     jiraConf,
		Sections: sections,
	}
	err = tmpl.Execute(&buf, &items)
	if err != nil {
		log.Fatalf("Error while applying template, the error is: %v", err)
	}
	log.Println(buf.String())
	return buf.String()
}
