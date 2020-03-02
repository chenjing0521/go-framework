package MiddleWares

import (
	"exercise/go-framework/Context"
	"exercise/go-framework/Route"
)

type MiddlewareFunc func(Route.HandleFunc) Route.HandleFunc

// MiddlewareLoggerFunc 函数实现日志中间件函数。
func MiddlewareLoggerFunc(h Route.HandleFunc) Route.HandleFunc {
	return func(ctx *Context.Context) {
		ctx.Printf("%s %s %s", ctx.RemoteAddr(), ctx.Method(), ctx.Path())
		h(ctx)
	}
}
