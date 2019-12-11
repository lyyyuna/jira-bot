package internal

type TomlConfig struct {
	Host     string
	Project  string
	Wxkey    string
	Username string
	Password string
	Users    map[string]JiraUser
	Filter   []JiraFilter
}

type JiraUser struct {
	Name string
}

type JiraFilter struct {
	Name  string
	Jql   string
	Split bool
}
