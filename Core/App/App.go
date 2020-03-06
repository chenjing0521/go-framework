package App

import (
	"exercise/go-framework/Core/Context"
	"exercise/go-framework/Core/Log"
	"exercise/go-framework/Core/MiddleWares"
	"exercise/go-framework/Core/Route"
	"net/http"
)

type Server interface {
	ListenAndServe() error
}

// App 框架主体
type App struct {
	Addr string
	Log.Logger
	Route.Router
	Server
	Middlewares []MiddleWares.MiddlewareFunc
}

// ServeHTTP 方式实现http.Hander接口，接管处理Http请求。
func (app *App) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := &Context.Context{
		Request:        req,
		ResponseWriter: resp,
		Logger:         app,
	}
	// 路由匹配
	h := app.Router.Match(ctx.Method(), ctx.Path())
	// 处理中间件
	for _, i := range app.Middlewares {
		h = i(h)
	}
	// 处理请求
	h(ctx)
}

// AddMiddleware App增加一个处理中间件。
func (app *App) AddMiddleware(m ...MiddleWares.MiddlewareFunc) {
	app.Middlewares = append(app.Middlewares, m...)
}

// Run 方法启动App。
func (app *App) Run() error {
	// Server初始化
	if app.Server == nil {
		app.Server = &http.Server{
			Addr:    app.Addr,
			Handler: app,
		}
	}
	app.Printf("start server: %s", app.Addr)
	return app.Server.ListenAndServe()
}

// NewApp 函数创建一个app。
func NewApp() *App {
	return &App{
		Addr:   ":8088",
		Logger: &Log.MyLogger{},
		Router: &Route.MyRouter{},
	}
}
