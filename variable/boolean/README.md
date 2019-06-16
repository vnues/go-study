# 布尔类型

## 基本介绍
- 布尔类型也叫 bool 类型，bool 类型数据只允许取值 true 和 false
- bool类型占1个字节
- bool 类型适于逻辑运算，一般用于程序流程控制[注:这个后面会详细介绍]
1.if条件控制语句
2.for循环控制语句

```go
package main
import (
	"fmt"
	"unsafe"
)

//演示golang中bool类型使用
func main() {
	var b = false
	fmt.Println("b=", b)
	//注意事项
	//1. bool类型占用存储空间是1个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b) )
	//2. bool类型只能取true或者false
	
}
```