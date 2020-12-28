package chat

type Chat struct {
	Id   int
	Name string
}

type Message struct {
	Id        int
	ChatId    int
	Body      string
	Timestamp int64
}
