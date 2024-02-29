

### go command
```shell
# 七牛 CDN 修改全局下载源
go env -w  GOPROXY=https://goproxy.cn,direct

# 初始化程序包 go.mod
go mod init <package>

# 检查缺失的依赖并安装 go.sum
go mod tidy

# 修改模块 （可以修改依赖地址）
go mod edit  -replace greetings=../greetings

```
