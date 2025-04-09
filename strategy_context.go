package fxtrategy

import (
	"go.uber.org/fx"
	"log"
)

// Context to inject and use Strategies
type Context[T any] struct {
	mapping map[string]T
}

type contextIn[T any] struct {
	fx.In
	Strategies []NamedStrategy[T] `group:"fxtrategy_common"`
}

func NewContext[T any](in contextIn[T]) *Context[T] {
	log.Println(in.Strategies)
	strategies := make(map[string]T, len(in.Strategies))
	for _, strategy := range in.Strategies {
		strategies[strategy.Name] = strategy.Item
	}
	return &Context[T]{
		mapping: strategies,
	}
}

func (c *Context[T]) Get(name string) (T, bool) {
	t, ok := c.mapping[name]
	return t, ok
}
