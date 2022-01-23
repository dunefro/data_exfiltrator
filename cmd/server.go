/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/dunefro/data_exfiltrator/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "creating server",
	Long:  `This will create a server at a specified port for connection and output to directed file`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("output")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		err := server.Serve(fileName, host, port)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// defining flags
	serverCmd.PersistentFlags().StringP("output", "o", "", "output(text file) to transfer the data (required)")
	serverCmd.MarkPersistentFlagRequired("output")
	serverCmd.PersistentFlags().StringP("host", "", "127.0.0.1", "host that you wish to connect")
	serverCmd.PersistentFlags().StringP("port", "p", "8080", "port that you wish to connect")
}
