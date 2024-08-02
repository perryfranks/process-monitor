package models

type Process struct {
	Name      string `json:"procName"`
	Workspace string `json:"workspaceName"`
	Id        int    `json:"ID"`
	IdString  string // Just for templ
}
