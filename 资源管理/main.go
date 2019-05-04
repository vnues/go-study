package main

import "fmt"

func main() {
	//response, err := http.Get("http://www.baidu.com")
	//if err != nil {
	//	// handle error
	//}
	////程序在使用完回复后必须关闭回复的主体。
	//defer response.Body.Close()
	//
	//body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	panic("error")
}

//函数执行完毕才会执行defer 有panic是先执行pannic之前的
//recover会捕获panic 通常写在defer里
//Recover捕获异常
//如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，
//并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。