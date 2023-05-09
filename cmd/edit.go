/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")

		targetPattern, err := cmd.Flags().GetString("target-pattern")
		if err != nil {
		}
		hostName, err := cmd.Flags().GetString("hostname")
		if err != nil {
		}

		fmt.Println(targetPattern, hostName)

	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("target-pattern", "T", "", "ip address")

	editCmd.Flags().StringP("hostname", "n", "", "port number")

	editCmd.MarkFlagRequired("target-pattern")
}
