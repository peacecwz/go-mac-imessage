package cmd

import (
	"fmt"
	"github.com/peacecwz/go-mac-imessage/sms"
	"github.com/spf13/cobra"
)

var message, to string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message",
	Long:  `Send a message`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := sms.Send(message, to)
		if err == nil {
			fmt.Println("Sent")
		}

		return err
	},
}

func init() {
	sendCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Message to send")
	sendCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "Phone number to send")

	RootCmd.AddCommand(sendCmd)
}
