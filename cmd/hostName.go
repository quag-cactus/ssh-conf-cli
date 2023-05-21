/*
Copyright © 2023 quag-cactus <quag.cactus@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kevinburke/ssh_config"
	"github.com/spf13/cobra"

	"github.com/quag-cactus/ssh-conf-cli/proc"
	"github.com/quag-cactus/ssh-conf-cli/utils"
)

// hostNameCmd represents the hostName command
var hostNameCmd = &cobra.Command{
	Use:   "hostName",
	Short: "Edit ssh_config value 'HostName' of designated Host",
	Long: `This subcommand edits the value of HostName for the host specified in the parent command.
	ssh-conf-cli edit -T <YourSpecifiedHost> hostName -n <aaa.bbb.ccc.ddd>`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("HostName rewriting...")

		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
		}
		targetPattern, err := cmd.Flags().GetString("target-pattern")
		if err != nil {
		}
		hostName, err := cmd.Flags().GetString("hostname")
		if err != nil {
		}

		// If filePath is empty, set default config path based on runntime.OS
		if filePath == "" {
			filePath = utils.DefineDefaultConfigPath()
			fmt.Printf("Config file path has been automatically defined: %s\n", filePath)
		} else {
			fmt.Printf("Config file path has been designated: %s\n", filePath)
		}

		// Create backup file
		backupFilePath, err := utils.CreateBackupConfigFile(filePath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to create backup file", backupFilePath)
			return
		} else {
			fmt.Println("Backup file has been created: ", backupFilePath)
		}

		inputFs, err := os.Open(filePath)
		inputFs.Seek(0, 0)
		cfg, err := ssh_config.Decode(inputFs)
		inputFs.Close()

		//fmt.Println()
		resultList, err := proc.RewriteConfigValue(cfg, targetPattern, "HostName", hostName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Searching for entries matching the host '%s'...\n", targetPattern)

		// if resultList is not empty, write to file
		if len(resultList) > 0 {
			fmt.Printf("Found %d entries\n", len(resultList))
			for i, result := range resultList {
				fmt.Printf("[%d] Host: %s\n", i+1, result.HostPatterns)
				fmt.Printf("[%d] Hostname is rewrited: %s -> %s (ln: %d)\n",
					i+1, result.PreviousValue, result.CurrentValue, result.LineNo)
			}
			utils.WriteConfigFile(filePath, cfg.String())
		} else {
			fmt.Printf("No entries found matching the host '%s'. Check target pattern and your config file.\n", targetPattern)
		}
	},
}

func init() {
	editCmd.AddCommand(hostNameCmd)
	hostNameCmd.Flags().StringP("hostname", "n", "", "HostName (ip address) for rewriting config")
}
