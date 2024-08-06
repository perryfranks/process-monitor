package models

import (
	"time"
)

type Process struct {
	Name      string `json:"procName"`
	Workspace string `json:"workspaceName"`
	User      string `json:"user"`
	Id        int    `json:"ID"`
	// May not always be available. Usually if output is captured process is blocked.
	// This may be overcome-able
	Pid         string    `json:"PID"`
	StartTime   time.Time `json:"startTime"`
	FinishTime  time.Time `json:"finishTime"`
	IdString    string    // Just for templ
	Finished    bool
	CapturedOut string
}
