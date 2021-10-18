package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("https://juejin.cn/")
	if err != nil {
		println(err)
	}
	// defer主要应用场景有异常处理、记录日志、清理数据、释放资源
	// 客户端关闭response的body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code:", resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// 不会输出Status Code accept-ranges等头部请求信息
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		println(err)
	}
	// 以字符串形式输出
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
