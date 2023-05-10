/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		targetPattern, err := cmd.Flags().GetString("target-pattern")
		if err != nil {
		}
		hostName, err := cmd.Flags().GetString("hostname")
		if err != nil {
		}

		fmt.Println("hostName called", targetPattern, hostName)
	},
}

func init() {
	editCmd.AddCommand(hostNameCmd)
	hostNameCmd.Flags().StringP("hostname", "n", "", "Host name (ip address) for rewriting config")
}
