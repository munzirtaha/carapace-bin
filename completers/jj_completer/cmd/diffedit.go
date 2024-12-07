package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var diffeditCmd = &cobra.Command{
	Use:   "diffedit",
	Short: "Touch up the content changes in a revision with a diff editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffeditCmd).Standalone()

	diffeditCmd.Flags().StringP("from", "f", "@", "Show changes from this revision. Defaults to @ if --to is specified")
	diffeditCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffeditCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	diffeditCmd.Flags().StringP("revision", "r", "@", "The revision to touch up. Defaults to @ if neither --to nor --from are specified")
	diffeditCmd.Flags().StringP("to", "t", "@", "Edit changes in this revision. Defaults to @ if --from is specified")
	rootCmd.AddCommand(diffeditCmd)

	carapace.Gen(diffeditCmd).FlagCompletion(carapace.ActionMap{
		"from":     jj.ActionRevs(jj.RevOption{}.Default()),
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":       jj.ActionRevs(jj.RevOption{}.Default()),
	})
}
