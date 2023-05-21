/*
Copyright © 2023 quag-cactus <quag.cactus@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configration value of ssh_config file",
	Long:  `This command edits the value of the key specified by the subcommand in the key-value pair of the specified host pattern.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.PersistentFlags().StringP("target-pattern", "T", "", "Target pattern of Host")

	editCmd.MarkPersistentFlagRequired("target-pattern")
}
