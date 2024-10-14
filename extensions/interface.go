package extensions

type IPublisher interface {
	Publish() error
}

type IConsumer interface {
	Consume() error
}
