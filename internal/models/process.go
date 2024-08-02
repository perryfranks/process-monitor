package models

type Process struct {
	Name      string `json:"procName"`
	Workspace string `json:"workspaceName"`
	Id        string `json:"ID"`
}
