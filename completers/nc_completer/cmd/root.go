package cmd

import (
	"github.com/rsteube/carapace-bin/completers/netcat_completer/cmd"
)

/**
Description for go:generate
	Short: "simple Unix utility which reads and writes data across network connections",
*/

func Execute() error {
	return cmd.ExecuteNc()
}