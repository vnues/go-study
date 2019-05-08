package main

import (
	"./filelisting"
	"fmt"
	"net/http"
	"os"
)
//我们不能把服务器报错的信息直接返回给用户看 而是得做一层错误处理函数
//字符串那块也很重要 rune
/**

熟悉统一错误处理的方式
error 和 panic ，能自己识别到的错误就用error，panic是一种碰到错误终止程序的错误，能不用尽量不用
一般使用 defer + panic + recover 结合进行统一错误处理
使用 Type Assertion 判断错误类型
结合函数式编程实现

**/

/**
slice的结构是这样子的
type slice struct {

	array unsafe.Pointer
	len   int
	cap   int
}

 **/

//handle业务逻辑处理--->做一层错误处理封装
//一个函数有返回值怎么使用？
//所以我们得传入这个函数参数进去使用
//函数当做参数需要声明类型 强类型语言的要求
//但是有个问题我直接用包对形式引用不就行了吗

//声明这种类型
type appHandler  func(writer http.ResponseWriter,request *http.Request) error

//什么时候需要定义函数类型 当它是作为参数时候
//对业务逻辑函数做错误处理 返回对是一个函数（这个函数又有返回值怎么写？）
func errWRapper(handle appHandler) func(writer http.ResponseWriter,request *http.Request) {
	//返回一个参数是writer http.ResponseWriter,request *http.Request的函数
	return func(writer http.ResponseWriter,request *http.Request){
		err := handle(writer,request)
		if err !=nil{
			//400
			code :=http.StatusOK
			switch{
			case os.IsNotExist(err):
				//404
				code=http.StatusNotFound
			case  os.IsPermission(err):
				//403权限
				code=http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}

	}

}



func main() {
	//字符串数组吗？
	//在 Go 中，一个字符串实际上就是一个只读的字节切片
	//本质是[]byte所以才可以循环
	var s string
	s="hello world"
	fmt.Println(s)
	fmt.Println(s[1]) //unicode编码

	s=s[2:4]
	fmt.Println(s)
	// ll
	//字符串也可以当slice使用？
	//执行这个步骤 会给回调函数传入request response
	 http.HandleFunc("/list/",errWRapper(filelisting.Handle))

	 //开启一个应用
	 err :=http.ListenAndServe(":8888",nil)
	 if err !=nil{
	 	panic(err)
	 }

}
//我希望现在的你 不是每天想着逃离公司的业务能拖一天就是一天 而是想着如何如何快速完成它 有时间完成自己的事情