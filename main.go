package main

import (
	"fmt"
	"net/http"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!", http.StatusOK, quote.Go())
}

// go mod init <package>	初始化程序包 go.mod

// go mod tidy    检查缺失的依赖并安装 go.sum
