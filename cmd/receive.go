package cmd

import (
	"fmt"
	"github.com/peacecwz/pinsms/internal"
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

		sqlite := internal.NewSqlite()
		err := internal.TrackSMS(sqlite, interval, func(sms []internal.SMS) {
			for _, s := range sms {
				fmt.Printf("Message: %s from %s\n", s.Content, s.From)
				err := sqlite.SetRead(s.Id)
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
