package cmd

import (
	"fmt"
	"github.com/peacecwz/mac-sms-tracker/internal"
	"github.com/spf13/cobra"
	"log"
)

var intervalSecond int64

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive a message",
	Long:  `Receive a message`,
	RunE: func(cmd *cobra.Command, args []string) error {

		interval, _ := cmd.Flags().GetInt64("interval")

		err := internal.TrackSMS(interval, func(sms []internal.SMS) {
			for _, s := range sms {
				fmt.Printf("Message: %s from %s\n", s.Content, s.From)
				err := s.Read()
				if err != nil {
					log.Fatalln(err)
				}
			}
		})
		return err
	},
}

func InitializeReceiveCmd(cmd *cobra.Command) {
	receiveCmd.PersistentFlags().Int64VarP(&intervalSecond, "interval", "i", 100, "Interval to check for new messages. It's millisecond")
	cmd.AddCommand(receiveCmd)
}
