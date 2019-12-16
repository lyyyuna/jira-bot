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
## {{.Section.Name}}
总数： [{{.Section.Cnt}}]({{.Section.Url}})
{{if .Section.Split}}
{{range $k, $v := .Section.Users}}{{$k}} 有 [{{$v.Cnt}}]({{$v.Url}})
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
	c := NewWechatClient(jiraConf.Wxkey)

	for _, section := range sections {
		content := markdown(section)
		c.SendToWechat(content)
	}

}

func markdown(section *SectionStats) string {
	tmpl, err := template.New("jira").Parse(jiraTemplate)
	if err != nil {
		log.Fatalf("Error while parsing template, the error is: %v", err)
	}

	var buf bytes.Buffer
	items := struct {
		Info    TomlConfig
		Section *SectionStats
	}{
		Info:    jiraConf,
		Section: section,
	}
	err = tmpl.Execute(&buf, &items)
	if err != nil {
		log.Fatalf("Error while applying template, the error is: %v", err)
	}
	log.Println(buf.String())
	return buf.String()
}
