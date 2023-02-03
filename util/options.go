package util

type Options struct {
	Verbose    bool
	ShowStatus bool
	Project    Project
}

type Project struct {
	ID   string
	Name string
}
