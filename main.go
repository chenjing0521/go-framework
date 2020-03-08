package main

import (
	"exercise/go-framework/Core/App"
	"exercise/go-framework/Core/Context"
	"exercise/go-framework/Core/MiddleWares"
	"fmt"
)

func main() {
	app := App.NewApp()
	app.AddMiddleware(MiddleWares.MiddlewareLoggerFunc)
	app.RegisterFunc("GET", "/hello", func(ctx *Context.Context) {
		fmt.Println(123123123123123)
		ctx.WriteString("hello micro web")
	})
	app.Run()
}
