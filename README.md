# Learning Go Programming language

---

## Go command:

go build：编译go源代码，在当前目录生成可执行文件（main package），非main package什么都不生成；

go install：编译go源代码，如果是main package会在GOPATH/bin生成可执行文件，非main package会在GOPATH/pkg生成packageName.a文件；

go get：git clone + go install；

go run：编译go main package源代码，并且执行，什么文件都不生成。
