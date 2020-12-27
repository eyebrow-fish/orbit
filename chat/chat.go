package chat

type Chat struct {
	Id   int
	Name string
}

type Message struct {
	ChatId    int
	AuthorId  int
	Timestamp int64
	Body      string
}
