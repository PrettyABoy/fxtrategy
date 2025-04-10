package fxtrategy

import (
	"go.uber.org/fx"
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

type TravelFunc[T any] func(string, T) error

// ForEach execute fn on each item
func (c *Context[T]) ForEach(fn TravelFunc[T]) error {
	for k, v := range c.mapping {
		err := fn(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Names returns all item's name
func (c *Context[T]) Names() []string {
	names := make([]string, 0, len(c.mapping))
	for name := range c.mapping {
		names = append(names, name)
	}
	return names
}
