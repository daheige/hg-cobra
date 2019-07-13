# cobra使用
    将命令cmd和入口文件main.go分开,可以用于多个项目规划，特别是涉及到rpc,web,job等

# 开始运行
    go mod tidy
    $ go run main.go 
    Using config file: /web/go/hg-cobra/app.yaml
    []
    name:  daheige
    OK
    author:  heige
    userInfo &{1 daheige}

    $ go run main.go version
    Using config file: /web/go/hg-cobra/app.yaml
    Demo Version: v1.0

    $ go run main.go test -n heige -a 123
    Using config file: /web/go/hg-cobra/app.yaml
    []
    name heige
    test called
    yaml name:  daheige
