## 资源管理及错误处理


### defer的应用(资源管理)

> 一个程序被打开那么就需要被关闭，这就是资源管理，在go语言中通常用defer实现


- defer 在函数结束时发生调用
- defer 的调用是栈类型 - 先进后出
- defer 通常用于资源关闭 Open/Close，Lock/UnLock 等


**一句话总结：defer 的调用机制是 “将defer语句加入栈中，当函数结束时（包括正常执行结束 / return / panic 出错结束等），从栈中依次执行 defer”**


举个例子：
```go
//这是一个错误处理函数
func writeFile(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close() // 将 "file.Close()" 压入 defer 栈中

    writer := bufio.NewWriter(file)
    defer writer.Flush() // 将 "writer.Flush()" 压入 defer 栈中

    fmt.Fprintln(writer, "123")
    // 当函数执行结束时，从 defer 栈中执行语句 - 后进先出，先 "writer.Flush()"，再 "file.Close()"
}

func main() {
    writeFile("defer.txt")
}
```

## 简单错误处理(通过注释处理)
   使用机制
   
   - 通过被调用函数的注释查看其可能发生的错误，然后依据错误类型并进行处理；
   - 错误处理结束后要 return




### panic&&recover(错误处理)

**一般不希望直接发生panic，所以尝试用recover进行阻止，recover能够捕获到panic**

panic

- 停止当前程序运行
- 一直向上返回，执行每一层的 defer
- 如果没有遇见 recover，程序退出

recover（相当于对 panic 的 catch 语句）

- 仅在 **defer** 调用中使用
- 获取 panic 的值
- 如果无法处理，可重新 panic

```go
 import (
     "fmt"
     "errors"
 )
 
 func recove() {
 	//最后函数退出才发生的
     defer func() {
         // func recover() interface{}，表示 recover() 函数的返回类型可以是各种类型，所以要判断是否是 error
         // 使用 recover() catch panic，防止程序直接退出
         r := recover()
         if err, ok := r.(error); ok {
             fmt.Println(err) // runtime error: integer divide by zero
         } else {
             panic(errors.New("not known error"))
         }
     }()
 
     b := 0
     a := 5/b // panic: runtime error: integer divide by zero
     fmt.Println(a)
 
     //panic("123") // panic: not known error
 }
 
 func main() {
     recove()
 }
```


## 错误统一处理

 >一个生产系统通常包含两种异常

- 不可直接暴露给用户的异常：例如系统内部异常
- 需要暴露给用户的异常：例如部分自定义异常信息用于提示用户操作

综合上述所学的知识，怎么实现一个错误统一处理（错误统一处理容易跟资源管理混错）

![](./filelisting/err.png)

> handle.go是业务逻辑部分（比如返回页面信息给用户），我们只管return错误就行，我们用errWRapper函数处理错误信息

```go
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

```
**errWRapper是个错误封装函数，我们对handle函数传过来对值作为一层处理，此处handle函数作为参数传入，所以我们得定义handle函数的类型**



> userError 自定义用户异常接口 此处就是定义给·用户看还是给系统看的错误处理
```go
type userError interface {
	error //内嵌类型
	//为什么一定需要这个类型 因为统一处理错误函数是返回error类型的--你只要实现这个方法那么可以理解成就是这个类型
	//duck typing
	Message() string
}
```


> 开启一个server服务器

```go
 http.HandleFunc("/",errWRapper(filelisting.Handle))

	 //开启一个应用
	 err :=http.ListenAndServe(":8888",nil)
	 if err !=nil{
	 	panic(err)
	 }
```


### 总结
- 能不用panic就不用panic
- 开启一个资源记得一定要关闭
- 类型断言可以作为条件判断，其中结构体.(接口)或者接口.(结构体)
- 一个函数作为参数就得定义他的类型--强语言的特点