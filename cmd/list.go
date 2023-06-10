/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/quag-cactus/ssh-conf-cli/proc"
	"github.com/quag-cactus/ssh-conf-cli/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List host-patterns configured in ssh_config",
	Long: `This command lists host-patterns configured in ssh_config.

	If you want to list host-patterns recursively using <Include> directive, use -r option.

	Default entry ssh_config file path is ~/.ssh/config (linux and Mac) or %USERPROFILE%\.ssh\config (Windows).
	If you want to designate another file, use -f option.
	`,
	Example: `
	# Example for listing host-patterns configured in ssh_config:
	$ ssh-conf-cli list -r
	`,
	Run: func(cmd *cobra.Command, args []string) {

		confPath, err := cmd.Flags().GetString("file")
		if err != nil {
		}
		recursive, err := cmd.Flags().GetBool("recursive")
		if err != nil {
		}

		// If confPath is empty, set default config path based on runntime.OS
		if confPath == "" {
			confPath = utils.DefineDefaultConfigPath()
			fmt.Printf("Config file path has been automatically defined: %s\n", confPath)
		} else {
			fmt.Printf("Config file path has been designated: %s\n", confPath)
		}

		// Get host-pattern list
		fmt.Println("following host-patterns are configured:")
		resultList, err := proc.GetHostPtnList(confPath, recursive)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, hostPtn := range resultList {
			fmt.Printf("[%d: %s] configured by %s\n", i, hostPtn[0], hostPtn[1])
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("recursive", "r", false, "flag for recursive search and display host-patterns")
}
