package mock

// 类A

// Retriever TODO
type Retriever struct {
	Contents string
}

// Get TODO
func (r Retriever) Get(url string) string {

	return url
}

// Post TODO
func (r Retriever) Post(url string) string {

	return url
}

// 这种就是多态 反正我这个给出这个方法 但是方法里面的行为不同

// TestRetriever TODO
// 类B
type TestRetriever struct {
	Contents string
}

// Get TODO
func (r TestRetriever) Get(url string) string {

	return url + "123"
}

// Post TODO
func (r TestRetriever) Post(url string) string {

	return url
}
