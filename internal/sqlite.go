package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Sqlite struct {
	db     *sql.DB
	cursor int64
}

func NewSqlite() *Sqlite {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dbSource := fmt.Sprintf("%s/Library/Messages/chat.db", homeDir)
	db, err := sql.Open("sqlite3", dbSource)

	if err != nil {
		log.Fatal(err)
	}

	return &Sqlite{
		db: db,
	}
}

func (s *Sqlite) SetRead(id string) error {
	_, err := s.db.Exec(`UPDATE message SET is_read = 1 WHERE guid = ?;`, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *Sqlite) GetAllSMS() ([]SMS, error) {
	rows, err := s.db.Query(`
SELECT message.guid                                                                    as 'id',
       datetime(message.date / 1000000000 + strftime("%s", "2001-01-01"), "unixepoch") AS 'date',
       message.text                                                                    as 'content',
       message.is_from_me                                                              as 'isFromMe',
       chat.chat_identifier                                                            as 'from',
       message.service                                                                 as 'service',
       message.is_read                                                                 as 'isRead'
FROM chat
         JOIN chat_message_join ON chat."ROWID" = chat_message_join.chat_id
         JOIN message ON chat_message_join.message_id = message."ROWID"
WHERE message.date > ? AND message.is_read = 0
ORDER BY message_date ASC;
`, s.cursor)

	if err != nil {
		log.Fatal(err)
	}

	var messages []SMS
	for rows.Next() {
		var (
			id       string
			date     string
			content  string
			isFromMe bool
			from     string
			service  string
			isRead   bool
		)

		err = rows.Scan(&id, &date, &content, &isFromMe, &from, &service, &isRead)

		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, SMS{
			Id:       id,
			Date:     date,
			Content:  content,
			IsFromMe: isFromMe,
			From:     from,
			Service:  service,
			IsRead:   isRead,
		})
	}

	return messages, nil
}

func (s *Sqlite) Close() {
	err := s.db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}
