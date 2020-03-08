package main

import (
	"fmt"
	"go-framework/Core/App"
	"go-framework/Core/Context"
	"go-framework/Core/MiddleWares"
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
