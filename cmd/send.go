package cmd

import (
	"fmt"
	"github.com/peacecwz/go-mac-imessage/internal"
	"github.com/spf13/cobra"
)

var message, to string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message",
	Long:  `Send a message`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.Send(message, to)
		if err == nil {
			fmt.Println("Sent")
		}

		return err
	},
}

func InitializeSmsCmd(cmd *cobra.Command) {
	sendCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Message to send")
	sendCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "Phone number to send")

	cmd.AddCommand(sendCmd)
}
