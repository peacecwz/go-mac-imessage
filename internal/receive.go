package internal

import (
	"log"
	"time"
)

func TrackSMS(sqlite *Sqlite, interval int64, trigger func(sms []SMS)) error {
	for {
		sms, err := sqlite.GetAllSMS()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		trigger(sms)

		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
