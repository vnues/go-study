package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	// 需要绝对路径 不支持相对路劲
     var str string = "/Users/vnues/2019go学习/file/abc.txt"
     file,err:=os.OpenFile(str,os.O_RDONLY,0777)
     if err!=nil{
     	fmt.Println("打开文件发生错误:%v",err)
     	return
	 }
     defer file.Close()
     reader:= bufio.NewReader(file)
     // fmt.Println(reader)
     // 也就是说当你读到以换行符结束时候那么我们当下次继续读取它就从下次的位置开始
	 // fmt.Println(reader.ReadString('.')) // hello,world!...读出了全部 因为就是没有小数点 然后for无限循环会打印出EOF

     fmt.Println(reader.ReadString('\n')) // 读到换行符就结束 打印出来包括str和error->此处为nil
    for{
    	// ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。如果ReadString方法在读取到delim之前遇到了错误，
    	// 它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
		str,err:=reader.ReadString('\n')
		if err==io.EOF{
			break;
		}
		fmt.Println(str)
     }
     // fmt.Printf("%s",reader)
}


// open和openFile的区别

// Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。--也就是是只读模式

// Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。

// OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。


// copy 将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。
//
// 对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。
// 我的理解错了被str,err:=reader.ReadString('\n')带偏这个是读取自定义以'\n'为结束 然后就是下次读取会继续截取的位置
// 而他跟copy没什么练习的 copy是以读取到末尾再也读不到就是给报EOF就是copy结束

// 现在在公司的学习就是可能看一下就上手实战 但也是仅限于框架 基础的语言我需要打扎实的基础

// 带缓存的可以操作大文件

// 还有一点是路径是绝对路径
// 后面感觉得天天加班吗？这就有点恐怖了
// 快速成长吧
// 互联网还是有很多机会留给我们寒门贵子的
// 我选择拿出大部分的时间来学习后端不亏 坚持下去 后面就是关键了