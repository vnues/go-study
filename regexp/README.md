## go的正则

### 正则字符：
 ```go
   /*
    代码	说明
    	.	匹配除换行符以外的任意字符
    	\w	匹配字母或数字或下划线或汉字
    	\s	匹配任意的空白符
    	\d	匹配数字
    	\b	匹配单词的开始或结束
    	^	匹配字符串的开始
    	$	匹配字符串的结束
        [字符类]  匹配“字符类”中的一个字符，“字符类”见后面的说明
        [^字符类]   匹配“字符类”外的一个字符，“字符类”见后面的说明
   */
   ```
### 正则方法：
  ```go
/*
regexp是公共方法 r是结构体
regexp.MatchString
regexp.Compile
r.MatchString
r.FindString
r.FindStringIndex
r.FindStringSubmatch 这个方法返回全局匹配的字符串和局部匹配的字符，比如
FindStringSubmatchIndex 和上面的方法一样，不同的是返回全局匹配和局部匹配的 起始索引和结束索引
r.FindAllString 个方法返回所有正则匹配的字符，不仅仅是第一个（全局搜索/匹配）
r.FindAllStringSubmatchIndex 这个方法返回所有全局匹配和局部匹配的字符串起始索引和结束索引
r.FindAllString  为这个方法提供一个正整数参数来限制匹配数量
regexp.MustCompile  当使用正则表达式来创建常量的时候，你可以使用`MustCompile` 因为`Compile`返回两个值
r.ReplaceAllString
r.ReplaceAllFunc
*/
```

### 写法：
- [里面的表达式是||或者不是&&] 里面放的值就是实际的值 不是正则的匹配符号那样的 比如你a-z表示a-z的范围，放.就表示点 而不是表示任意字符

```go
r, _ := regexp.Compile("p([a-z.]+)ch") //a-z范围或者有点.都符合
```
- 正则的ALL指的是全局搜索 方法没有all不是全局搜索 all全局搜索如果需要int参数 -1代表全部
```go
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string {
	//代表如果小于0返回全部匹配到的字段
	if n < 0 {
		n = len(s) + 1
	}
	var result [][]string
	re.allMatches(s, nil, n, func(match []int) {
		if result == nil {
			result = make([][]string, 0, startSize)
		}
		slice := make([]string, len(match)/2)
		for j := range slice {
			if match[2*j] >= 0 {
				slice[j] = s[match[2*j]:match[2*j+1]]
			}
		}
		result = append(result, slice)
	})
	return result
}
```

- 以Submatch字段都方法都都是返回包括子正则选项
```go
r, _ := regexp.Compile("p([a-z.]+)ch")
fmt.Println(r.FindAllStringSubmatch("peaddddchqwqweqwe hello world",-1))
```
