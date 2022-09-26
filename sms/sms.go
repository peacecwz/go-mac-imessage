package sms

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
	db       *sqlite
}

func (s SMS) Read() error {
	return s.db.setRead(s.Id)
}
