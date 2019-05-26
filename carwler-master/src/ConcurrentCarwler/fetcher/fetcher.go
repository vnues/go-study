package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)
//只发送值的
/*
func Tick(d Duration) <-chan Time {
	if d <= 0 {
		return nil
	}
	return NewTicker(d).C
}
*/
//rateLimter是个channel 它的意思等待10 * time.Millisecond才执行下面这段赋值吗？
//每个fetcher断续去请求？不是的 因为我们执行的是Fetch不是整个包
//100个woker会抢这个rateLimter --所以这个rateLimter每隔10 * time.Millisecond往这管道放东西
//这只是个引用
//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。
// 它会调整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源

//类似js的setTimeout
var rateLimter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimter
	reps, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer reps.Body.Close()

	if reps.StatusCode != http.StatusOK {
		// fmt.Println("Error: Status Code: ", reps.StatusCode)
		return nil, fmt.Errorf("wrong status Code: %d", reps.StatusCode)
	}

	e := determineEncoding(reps.Body)
	utf8Reader := transform.NewReader(reps.Body,
		e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
