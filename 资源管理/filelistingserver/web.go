package main

import (
	"./filelisting"
	"fmt"
	"log"
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




/*
增加业务逻辑-->需要增加用户和系统错误处理
为什么需要？
//我们想要自定义返回的错误信息
*/

type userError interface {
	error
	Message() string
}



//声明这种类型
type appHandler  func(writer http.ResponseWriter,request *http.Request) error


/*
当defer语句被执行时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时,defer后的函数才会被执行，
不论包含defer语句的函数是通过return正常结束，
还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。
//先进后出
*/

//什么时候需要定义函数类型 当它是作为参数时候
//对业务逻辑函数做错误处理 返回对是一个函数（这个函数又有返回值怎么写？）
func errWRapper(handle appHandler) func(writer http.ResponseWriter,request *http.Request) {
	//返回一个参数是writer http.ResponseWriter,request *http.Request的函数
	return func(writer http.ResponseWriter,request *http.Request){
		//panic
        defer func (){
        	//自定义recover
        	//如果发生未知错误->自定义处理函数
        	/*
        	func Error(w ResponseWriter, error string, code int) {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.WriteHeader(code)
				fmt.Fprintln(w, error)
			}
        	*/
        	if r:=recover();r!=nil{
        		log.Printf("Panic:%v",r)
        		http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
		err := handle(writer,request)
		//有返回错误的值 error
		if err !=nil{
			//打印日志
			log.Printf("Error occurred " + "handle request: %s",err.Error())
			//userError
			if r,ok :=err.(userError);ok{
				  http.Error(writer,r.Message(),http.StatusBadRequest)
				  //return终止这个函数运行的意思
				  return
			}
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
	 http.HandleFunc("/",errWRapper(filelisting.Handle))

	 //开启一个应用
	 err :=http.ListenAndServe(":8888",nil)
	 if err !=nil{
	 	panic(err)
	 }

}
//我希望现在的你 不是每天想着逃离公司的业务能拖一天就是一天 而是想着如何如何快速完成它 有时间完成自己的事情