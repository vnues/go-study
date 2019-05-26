package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)
//难题不要去怕，去正面接触攻克它就行

func main(){
	//http要加进去
  resp,err:=http.Get("https://www.zhenai.com/zhenghun")
  if err!=nil{
  	panic(err)
  }
  //客户端关闭reponse的body
  defer resp.Body.Close()
  if resp.StatusCode!=http.StatusOK{
  	fmt.Println("Error:status code:",resp.StatusCode)
  	return
  }
	//Body io.ReadCloser  源代码
	//http客户端请求返回的是是个io.ReaderCloser类型 需要用io去读以[]byte返回
	/*
	  func ReadAll(r io.Reader) ([]byte, error) {
	  	return readAll(r, bytes.MinRead)
	  }
	 body源码
     type body struct {
	    src          io.Reader
	    hdr          interface{}
	    r            *bufio.Reader
	    closing      bool
	    doEarlyClose bool
	    mu         sync.Mutex
	    sawEOF     bool
	    closed     bool
	    earlyClose bool
	    onHitEOF   func()
	}
	*/
	//有一个问题，http是输入输出形式(有请求过来 read读出来发送给他 有数据是存入的就写入进去)所以返回的有包含reader类型
	//万一是gbk格式
	//Go语言对获取的html进行转码
	//一.首先我们需要下载两个包："golang.org/x/text" ”golang.org/x/net/html"，
	// 前者用来对HTML进行转码，而后者则是猜测出当前获取的HTML是哪一种编码，辅助前者完成转码。
	//详情见Determineencoding函数的介绍，它是根据reader的前1024个字节来猜测编码格式的。
	e :=determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	//不会输出Status Codet accept-ranges等头部请求信息
	all,err := ioutil.ReadAll(utf8Reader)
    //fmt.Println(utf8Reader)
	//fmt.Println(resp.Body)
  if err!=nil{
  	panic(err)
  }
  //以字符串形式输出
  fmt.Printf("%s\n",all)
}

func determineEncoding(r io.Reader) encoding.Encoding{
      //// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	  bytes,err:= bufio.NewReader(r).Peek(1024)
	  if err!=nil{
	  	panic(err)
	  }
	  //func DetermineEncoding(content []byte, contentType string) (e encoding.Encoding, name string, certain bool)
	 //charset -->utf-8
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	  return e
}


//go被墙怎么解决
/*
github解决办法
其实 golang 在 github 上建立了一个镜像库，如 https://github.com/golang/net 即是 https://golang.org/x/net 的镜像库

获取 golang.org/x/net 包，其实只需要以下步骤：

在命令行里依次输入如下：


1.mkdir -p $GOPATH/src/golang.org/x
2.cd $GOPATH/src/golang.org/x
3.git clone https://github.com/golang/net.g
*/


//写代码要提高通用性
//gopm怎么放在全局下

//bufio 包提供了两个实例化 bufio.Reader 对象的函数：NewReader 和 NewReaderSize。其中，NewReader 函数是调用 NewReaderSize 函数实现的：
//bufio 用来帮助处理 I/O 缓存。 我们将通过一些示例来熟悉其为我们提供的：Reader, Writer and Scanner 等一系列功能
//多次进行小量的写操作会影响程序性能。每一次写操作最终都会体现为系统层调用，频繁进行该操作将有可能对 CPU 造成伤害。而且很多硬件设备更适合处理块对齐的数据，例如硬盘。为了减少进行多次写操作所需的开支，golang 提供了 bufio.Writer。数据将不再直接写入目的地(实现了 io.Writer 接口)，而是先写入缓存，当缓存写满后再统一写入目的地：



//为什么打印出来没有docype	标签
//%v 以默认的方式打印变量的值
//%T 打印变量的类型