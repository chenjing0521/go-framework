package Context

import (
	"exercise/go-framework/Core/Log"
	"net/http"
	"strings"
)

// Context 请求上下文，封装请求操作，未详细实现。
type Context struct {
	*http.Request
	http.ResponseWriter
	Log.Logger
}

// WriteString 方法实现请求上下文返回字符串。
func (ctx *Context) WriteString(s string) {
	ctx.ResponseWriter.Write([]byte(s))
}

// Method 方法获取请求方法。
func (ctx *Context) Method() string {
	return ctx.Request.Method
}

// Path 方法获取请求路径。
func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

// RemoteAddr 方法获取客户端真实地址。
func (ctx *Context) RemoteAddr() string {
	xforward := ctx.Request.Header.Get("X-Forwarded-For")
	if "" == xforward {
		return strings.SplitN(ctx.Request.RemoteAddr, ":", 2)[0]
	}
	return strings.SplitN(string(xforward), ",", 2)[0]
}
