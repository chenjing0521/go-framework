package main

import (
	"exercise/go-framework/App"
	"exercise/go-framework/Context"
	"exercise/go-framework/MiddleWares"
)

func main() {
	app := App.NewApp()
	app.AddMiddleware(MiddleWares.MiddlewareLoggerFunc)
	app.RegisterFunc("GET", "/hello", func(ctx *Context.Context) {
		ctx.WriteString("hello micro web")
	})
	app.Run()
}
