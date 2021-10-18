package fetch

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	_ "time" // time TODO

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 声明函数需要参数名 返回的=需要声明类型 error是全局类型
// <-channel类型
var rateLimter = time.Tick(40 * time.Millisecond)

// Fetcher 间断器利用的就是阻塞
func Fetcher(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 有请求但是不成功
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code :%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	/*
	   NewReader returns a new Reader whose buffer has the default size.
	   func NewReader(rd io.Reader) *Reader {
	   	return NewReaderSize(rd, defaultBufSize)
	   }
	*/
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 传的是指针但是不会改变--阅读源代码
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		// 返回默认的
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
