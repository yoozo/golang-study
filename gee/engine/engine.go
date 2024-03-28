package engine

import (
	"fmt"
	"net/http"
	"reflect"
)

type Engine struct {
	http.Handler
	appName   string
	routerMap map[string](map[string]http.HandlerFunc)
}

func New(appName string) *Engine {
	return &Engine{
		appName:   appName,
		routerMap: make(map[string](map[string]http.HandlerFunc)),
	}
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	// Method: POST
	// URL: /debug
	// Proto: HTTP/1.1
	// ProtoMajor: 1
	// ProtoMinor: 1
	// Header: map[Accept:[*/*] User-Agent:[curl/8.1.2]]
	// Body: {}
	// GetBody: <nil>
	// ContentLength: 0
	// TransferEncoding: []
	// Close: false
	// Host: localhost:8080
	// Form: map[]
	// PostForm: map[]
	// MultipartForm: <nil>
	// Trailer: map[]
	// RemoteAddr: 127.0.0.1:53278
	// RequestURI: /debug
	// TLS: <nil>
	// Cancel: <nil>
	// Response: <nil>

	// fmt.Println("app: ", engine.appName)
	// fmt.Println("res.URL: ", res.URL)
	// printStructKeysAndValues(res)

	method := res.Method
	path := res.URL.Path

	handlerFunc, ok := engine.routerMap[path][method]
	if !ok {
		fmt.Fprintf(w, "method %v %v is 404 Not Found", method, path)
		return
	}
	handlerFunc(w, res)
}

func (engine Engine) AddRouter(method string, urlPath string, f http.HandlerFunc) {
	_, ok := engine.routerMap[urlPath]
	if !ok {
		engine.routerMap[urlPath] = make(map[string]http.HandlerFunc)
	}
	engine.routerMap[urlPath][method] = f
}

func printStructKeysAndValues(s interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println("Provided type is not a struct")
		return
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("%s: %v\n", t.Field(i).Name, f.Interface())
	}
}
