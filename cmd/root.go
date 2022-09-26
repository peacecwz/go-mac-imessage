package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sms",
	Short: "sms is a command line tool for sending and receiving SMS messages",
	Long:  `sms is a command line tool for sending and receiving SMS messages`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
