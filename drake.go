package drake

type Session struct {
	ID string
}

type Database interface {
	CreateSession(*Session) error
}
