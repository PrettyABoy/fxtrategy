package fxtrategy_test

import (
	"fmt"
	"github.com/PrettyABoy/fxtrategy"
	"go.uber.org/fx"
	"testing"
)

func TestInjection(t *testing.T) {
	fx.New(
		fx.Provide(newA, newB),
		fx.Provide(fxtrategy.NewContext[Fool]),
		fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
			a, _ := ctx.Get("a")
			a.FoolSpeak()
			b, _ := ctx.Get("b")
			b.FoolSpeak()
		}),
	)
}

func TestMultiple(t *testing.T) {
	fx.New(
		fx.Provide(newA, newB, newC, newD),
		fx.Provide(fxtrategy.NewContext[Fool]),
		fx.Provide(fxtrategy.NewContext[Genius]),
		fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
			a, _ := ctx.Get("a")
			a.FoolSpeak()
			b, _ := ctx.Get("b")
			b.FoolSpeak()
			_, ok := ctx.Get("c")
			if ok {
				t.Error("c shouldn't be ok")
			}
		}),
		fx.Invoke(func(ctx *fxtrategy.Context[Genius]) {
			c, _ := ctx.Get("c")
			c.GeniusSpeak()
			d, _ := ctx.Get("d")
			d.GeniusSpeak()
		}),
	)
}

type Fool interface {
	FoolSpeak()
}

func newA() fxtrategy.Strategy[Fool] {
	return fxtrategy.Strategy[Fool]{
		NS: fxtrategy.NamedStrategy[Fool]{
			Name: "a",
			Item: &FoolA{},
		},
	}
}

func newB() fxtrategy.Strategy[Fool] {
	return fxtrategy.Strategy[Fool]{
		NS: fxtrategy.NamedStrategy[Fool]{
			Name: "b",
			Item: &FoolB{},
		},
	}
}

type FoolA struct {
}

func (a *FoolA) FoolSpeak() {
	fmt.Println("i'm a")
}

type FoolB struct {
}

func (b *FoolB) FoolSpeak() {
	fmt.Println("i'm b")
}

type Genius interface {
	GeniusSpeak()
}

func newC() fxtrategy.Strategy[Genius] {
	return fxtrategy.Strategy[Genius]{
		NS: fxtrategy.NamedStrategy[Genius]{
			Name: "c",
			Item: &GeniusC{},
		},
	}
}

func newD() fxtrategy.Strategy[Genius] {
	return fxtrategy.Strategy[Genius]{
		NS: fxtrategy.NamedStrategy[Genius]{
			Name: "d",
			Item: &GeniusD{},
		},
	}
}

type GeniusC struct {
}

func (c *GeniusC) GeniusSpeak() {
	fmt.Println("i'm c")
}

type GeniusD struct {
}

func (d *GeniusD) GeniusSpeak() {
	fmt.Println("i'm d")
}
