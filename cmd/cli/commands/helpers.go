package commands

import "strings"

// take a command as a string and then split it into words
func parseQuoteCmd(s string) []string {
	var invocation []string
	invocation = strings.Split(s, " ")

	return invocation

}
