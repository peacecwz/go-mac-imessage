package main

import (
	"fmt"
	"github.com/peacecwz/go-mac-imessage/sms"
	"time"
)

func main() {
	interval := int64(100)
	go sms.TrackSMS(interval, func(sms []sms.SMS) {
		for _, s := range sms {
			if s.IsRead {
				continue
			}
			s.Read()
			fmt.Printf("Message: %s From: %s\n", s.Content, s.From)
		}
	})

	sms.Send("Hello World", "+310638561717")

	time.Sleep(5 * time.Second)
}
