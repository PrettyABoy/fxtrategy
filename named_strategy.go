package fxtrategy

import "go.uber.org/fx"

// NamedStrategy warp Strategy with name and a default group
type NamedStrategy[T any] struct {
	Name string
	Item T
}

type Strategy[T any] struct {
	fx.Out
	NS NamedStrategy[T] `group:"fxtrategy_common"`
}
