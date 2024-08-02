package main

import (
	"html/template"
	"strings"
)

type DisplayVars struct {
	CurrentPage string
}

func (d *DisplayVars) NewDisplayVars() {
}

func isCurrentPageColor(queryPage string, d DisplayVars) template.HTMLAttr {
	queryPage = strings.TrimSpace(queryPage)
	if queryPage == d.CurrentPage {
		return "text-tOrange underline"
	}

	return ""
}
