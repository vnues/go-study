package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// Retriever 实现了这个接口的所有方法才说实现了这个接口
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// Get 方法是个值接收者可以接受值或者指针
// 但是指针接收者只能接收指针
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}
