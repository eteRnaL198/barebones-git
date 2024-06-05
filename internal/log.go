package internal

import "strings"

func PrintLog(logFileContents string) string {
	var logs []string
	lines := strings.Split(logFileContents, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		logs = append(logs, line)
	}
	return strings.Join(logs, "\n")
}
