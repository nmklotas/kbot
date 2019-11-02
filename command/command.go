package command

type Message struct {
	Text   string
	UserId string
}

type Command interface {
	CanHandle(message Message) bool
	Handle(message Message) error
	Help() string
}
