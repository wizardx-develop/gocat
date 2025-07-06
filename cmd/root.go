/*
Copyright © 2025 wizardx nexuslev@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"wizardx/gocat/concat"

	"github.com/spf13/cobra"
)

var showEnds bool
var showTabs bool
var showStrNumber bool
var showNonEmptyStrNumbers bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocat",
	Short: "Concatenate FILE(s) to standard output.",
	Long: `Concatenate FILE(s) to standard output.
	
With no FILE, or when FILE is -, read standard input.`,
	Run: func(cmd *cobra.Command, args []string) {
		var result string

		if len(args) == 0 {
			concat.InToOut()
		} else {
			for _, fname := range args {
				line := concat.Concat(fname)

				if showNonEmptyStrNumbers == true && showStrNumber == false {
					line = concat.ShowNonEmptyStrNumbers(line)
				}
				if showEnds == true {
					line = concat.ShowEnds(line)
				}
				if showTabs == true {
					line = concat.ShowTabs(line)
				}
				if showStrNumber == true {
					line = concat.ShowStrNumbers(line)
				}
				result += concat.ToString(line)

				fmt.Fprint(os.Stdout, result)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showEnds, "show-ends", "E", false, "display $ at end of each line")
	rootCmd.PersistentFlags().BoolVarP(&showTabs, "show-tabs", "T", false, "display TAB characters as ^I")
	rootCmd.PersistentFlags().BoolVarP(&showStrNumber, "number", "n", false, "number all output lines")
	rootCmd.PersistentFlags().BoolVarP(&showNonEmptyStrNumbers, "number-nonblank", "b", false, "number nonempty output lines, overrides -n")
}
