package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionFilters() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-filters")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^.{4} (?P<filter>[^ ]+) +[^ ]+ +(?P<description>.*)`)

		vals := make([]string, 0)
		for _, line := range lines[8 : len(lines)-1] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
