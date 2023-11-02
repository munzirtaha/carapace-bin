package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var interdiffCmd = &cobra.Command{
	Use:   "interdiff",
	Short: "Compare the changes of two commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(interdiffCmd).Standalone()

	interdiffCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	interdiffCmd.Flags().String("from", "", "Show changes from this revision")
	interdiffCmd.Flags().Bool("git", false, "Show a Git-format diff")
	interdiffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	interdiffCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	interdiffCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	interdiffCmd.Flags().String("to", "", "Show changes to this revision")
	interdiffCmd.Flags().String("tool", "", "Generate diff by external command")
	interdiffCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(interdiffCmd)
}