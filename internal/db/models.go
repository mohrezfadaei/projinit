package db

type License struct {
	ID      int
	Type    string
	Content string
}

type Gitignore struct {
	ID       int
	Language string
	Content  string
}
