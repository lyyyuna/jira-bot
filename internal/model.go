package internal

type SectionStats struct {
	Name  string
	Jql   string
	Url   string
	Cnt   int
	Users map[string]*UserStats
	Split bool
}

type UserStats struct {
	Jql string
	Cnt int
	Url string
}
