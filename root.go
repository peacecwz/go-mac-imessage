package main

import (
	"github.com/peacecwz/mac-sms-tracker/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "sms",
		Short: "sms is a command line tool for sending and receiving SMS messages",
		Long:  `sms is a command line tool for sending and receiving SMS messages`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.InitializeSmsCmd(rootCmd)
	cmd.InitializeReceiveCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
