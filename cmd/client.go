/*
Copyright © 2022 Vedant Pareek <pareekvedant99@gmail.com>

*/
package cmd

import (
	"log"

	"github.com/dunefro/data_exfiltrator/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "to run the client",
	Long:  `Running the client for data exfiltrator`,
	Run: func(cmd *cobra.Command, args []string) {

		fileName, _ := cmd.Flags().GetString("file")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		err := client.ExfiltrateFile(fileName, host, port)
		if err != nil {
			log.Println("Failed to transfer the file")
			log.Fatal(err.Error())
		} else {
			log.Println("Successful: File was transferred")
		}

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// defining flags for client
	clientCmd.PersistentFlags().StringP("file", "f", "", "file(text) name which you want to transfer (required)")
	clientCmd.MarkPersistentFlagRequired("file")
	clientCmd.PersistentFlags().StringP("host", "", "127.0.0.1", "host that you wish to connect")
	clientCmd.PersistentFlags().StringP("port", "p", "8080", "port that you wish to connect")

}
