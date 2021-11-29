# go-script
写一些go 的脚本

## gocheck
### 安装
```shell
git clone git@github.com:weirwei/go-script.git
```
```shell
cd go-script/gocheck
```
```shell
go install
```

使用前需先安装golangci-lint
```shell
# 安装golangci-lint
go get github.com/golangci/golangci-lint/cmd/golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```
### 用法
```shell
# 使用方法
gocheck -h

Usage of gocheck:
  -d string
        检查的项目目录 (default "./")
```

