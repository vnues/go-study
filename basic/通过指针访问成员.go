package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
//我的认知里就是通过拿到值就可以拿到成员 但是指针就可以了 
func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700
   
   var struct_pointer *Books= &Book1;
   fmt.Println("拿到的成员")
   fmt.Println(struct_pointer)
   //&{Go 语言 www.runoob.com Go 语言教程 6495407}
   fmt.Println(struct_pointer.title)
   fmt.Println(*struct_pointer)
  //打印出来的其实就是一串值{Go 语言 www.runoob.com Go 语言教程 6495407}
  // fmt.Println(*struct_pointer.title)
  
}

