package craw

var client Factory

type Factory interface {
	HuYaCraw() HuYaCraw
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
