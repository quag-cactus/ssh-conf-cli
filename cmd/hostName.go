/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

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
		}

		// Create backup file
		backupFilePath, err := utils.CreateBackupConfigFile(filePath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to create backup file", backupFilePath)
			return
		}

		inputFs, err := os.Open(filePath)
		inputFs.Seek(0, 0)
		cfg, _ := ssh_config.Decode(inputFs)
		inputFs.Close()

		for _, host := range cfg.Hosts {
			fmt.Println("patterns:", host.Patterns)
		}

		//xList, x, _ := proc.TestEdit()
		resultList, err := proc.RewriteConfigValue(cfg, targetPattern, "HostName", hostName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resultList)

		// if resultList is not empty, write to file
		if len(resultList) > 0 {
			utils.WriteConfigFile(filePath, cfg.String())
		}

		fmt.Println("hostName called", filePath, targetPattern, hostName)
	},
}

func init() {
	editCmd.AddCommand(hostNameCmd)
	hostNameCmd.Flags().StringP("hostname", "n", "", "Host name (ip address) for rewriting config")
}
