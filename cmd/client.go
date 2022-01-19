/*
Copyright © 2022 Vedant Pareek <pareekvedant99@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/dunefro/data_exfiltrator/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called from client cobra")
		fileName, _ := cmd.Flags().GetString("file")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		client.ExfiltrateFile(fileName, host, port)
		// err := client.CheckFile(fileName)
		// if err != nil {
		// 	fmt.Errorf("File %s doesn't exist.", fileName)
		// }

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	clientCmd.PersistentFlags().StringP("file", "f", "", "file(text) name which you want to transfer (required)")
	clientCmd.MarkPersistentFlagRequired("file")
	clientCmd.PersistentFlags().StringP("host", "", "127.0.0.1", "host that you wish to connect")
	clientCmd.PersistentFlags().StringP("port", "p", "8080", "port that you wish to connect")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
