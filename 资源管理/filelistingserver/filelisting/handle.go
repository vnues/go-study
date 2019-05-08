package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
//业务逻辑部分不做错误处理 有错误我就给返回没有错误就返回nil
func Handle(writer http.ResponseWriter,request *http.Request) error{
	fmt.Println(request.URL.Path)
	path :=request.URL.Path[len("/list/"):]
	fmt.Println(path)

	file,err := os.Open(path)
	if err !=nil{
		//http server如果发生panic会被保护起来 不被终止掉
		//写入这个wirter
		//http.Error(writer,err.Error(),
		//	http.StatusInternalServerError)
		return err
		//panic(err)
	}

	defer file.Close()

	all,err:=ioutil.ReadAll(file)
	if err !=nil{
		return err
		//panic(err)
	}
	//resopnse
	writer.Write(all)
	return nil
}