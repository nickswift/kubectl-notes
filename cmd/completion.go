package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:       "completion",
	Short:     "Generates auto-complete scripts",
	Args:      cobra.ExactArgs(1),
	Run:       completionCmdRun,
	ValidArgs: []string{"bash", "zsh"},
	PreRunE:   completionCmdPreRunE,
}

func completionCmdRun(_ *cobra.Command, args []string) {
	shell := args[0]
	switch shell {
	case "bash":
		exitOnErr(rootCmd.GenBashCompletion(os.Stdout))
	case "zsh":
		exitOnErr(rootCmd.GenZshCompletion(os.Stdout))
	}
}

func contains(sl []string, s string) bool {
	for _, a := range sl {
		if a == s {
			return true
		}
	}
	return false
}

func completionCmdPreRunE(cmd *cobra.Command, args []string) error {
	shell := args[0]
	validShells := cmd.ValidArgs

	if !contains(validShells, shell) {
		return errors.New(fmt.Sprintf("The shell %s is not recognized. Valid shells: %s", shell, validShells))
	}

	return nil
}
