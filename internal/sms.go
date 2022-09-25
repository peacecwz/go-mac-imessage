package internal

type SMSInterface interface {
	Read() error
}

type SMS struct {
	Id       string
	Date     string
	Content  string
	IsFromMe bool
	From     string
	Service  string
	IsRead   bool
	db       *Sqlite
}

func (s SMS) Read() error {
	return s.db.SetRead(s.Id)
}
