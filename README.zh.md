# Fxtrategy

[Eng](README.md) | 中文

## 📚 Table of Contents

- [简介](#简介)
- [特性](#特性)
- [GetStart](#get-start)
- [示例](#示例)

---

## 简介

`fxtrategy` 是一个基于 [Uber FX](https://github.com/uber-go/fx) 的策略注入工具，使用泛型与命名策略模式，为 Go 应用程序提供灵活、模块化的策略注册与动态获取能力。

---

## 特性

- ✅ 支持通过接口定义策略行为
- ✅ 基于泛型实现强类型上下文 Context
- ✅ 可按名称注册与获取策略实例
- ✅ 支持多个上下文类型共存（如：`Fool` 和 `Genius`）
- ✅ 与 Uber FX 完美集成

---

## Get Start
```shell
go get github.com/PrettyABoy/fxtrategy@latest
```

---

## 示例

### 构造
```go
func newA() fxtrategy.Strategy[Fool] {
	return fxtrategy.Strategy[Fool]{
		NS: fxtrategy.NamedStrategy[Fool]{
			Name: "a",
			Item: &FoolA{},
		},
	}
}
```

### 带依赖的构造
```go
func newA(dependencyA any, dependencyB any) fxtrategy.Strategy[Fool] {
	return fxtrategy.Strategy[Fool]{
		NS: fxtrategy.NamedStrategy[Fool]{
			Name: "a",
			Item: &FoolA{
				depA: dependencyA,
				depB: dependencyB,
            },
		},
	}
}
```

### 使用
```go
fx.New(
    fx.Provide(newA, newB, newC, newD),
    fx.Provide(fxtrategy.NewContext[Fool]),
    fx.Provide(fxtrategy.NewContext[Genius]),

    fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
        a, _ := ctx.Get("a")
        a.Speak() // 输出：i'm a
        b, _ := ctx.Get("b")
        b.Speak() // 输出：i'm b
    }),

    fx.Invoke(func(ctx *fxtrategy.Context[Genius]) {
        c, _ := ctx.Get("c")
        c.Speak() // 输出：i'm c
        d, _ := ctx.Get("d")
        d.Speak() // 输出：i'm d
    }),
)
```
