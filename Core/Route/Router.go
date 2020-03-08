package Route

import (
	"go-framework/Core/Context"
	"strings"
)

// HandleFunc 请求处理函数
type HandleFunc func(ctx *Context.Context)

type Router interface {
	Match(string, string) HandleFunc
	RegisterFunc(string, string, HandleFunc)
}

type MyRouter struct {
	RoutesConst map[string]HandleFunc
	RoutesPath  []string
	RoutesFunc  []HandleFunc
}

// Match 方法匹配一个Context的请求，实现Router接口。
func (r *MyRouter) Match(method, path string) HandleFunc {
	path = method + path
	if h, ok := r.RoutesConst[path]; ok {
		return h
	}
	for i, p := range r.RoutesPath {
		if strings.HasPrefix(path, p) {
			return r.RoutesFunc[i]
		}
	}
	return Handle404
}

func (r *MyRouter) RegisterFunc(method string, path string, handle HandleFunc) {
	if r.RoutesConst == nil {
		r.RoutesConst = make(map[string]HandleFunc)
	}
	path = method + path
	if path[len(path)-1] == '*' {
		r.RoutesPath = append(r.RoutesPath, path)
		r.RoutesFunc = append(r.RoutesFunc, handle)
	} else {
		r.RoutesConst[path] = handle
	}
}

// Handle404 函数定义处理404响应，没有找到对应的资源。
func Handle404(ctx *Context.Context) {
	ctx.ResponseWriter.WriteHeader(404)
	ctx.ResponseWriter.Write([]byte("404 Not Found"))
}
