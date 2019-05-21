package engine

//封装的Request这个对象
type Request struct {
	Url string
	//解析函数
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	//任意类型
	//城市
	Items []interface{}
}

func NilParser([] byte) ParseResult{
	//返回空的东西给它
	return  ParseResult{}
}
