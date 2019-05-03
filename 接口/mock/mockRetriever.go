package mock

//类A

type  Retriever struct {
     Contents string
}


func (r Retriever) Get(url string) string{

	return url
}

func (r Retriever) Post(url string) string{

	return url
}

//这种就是多态 反正我这个给出这个方法 但是方法里面的行为不同

//类B
type  TestRetriever struct {
	Contents string
}


func (r TestRetriever) Get(url string) string{

	return url+"123"
}

func (r TestRetriever) Post(url string) string{

	return url
}
