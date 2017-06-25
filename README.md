# Learning Go Programming language

## Go command:

**go build：** 编译go源代码，在当前目录生成可执行文件（main package），non-main package无效；

**go install：** 编译go源代码，如果是main package会在$GOPATH/bin生成可执行文件，non-main package会在$GOPATH/pkg/$GOOS_$GOARCH生成packageName.a文件；

**go get：** git clone + go install；

**go run：** 编译go main package源代码，并且执行，不生成可执行文件。
