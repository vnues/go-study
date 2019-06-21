package main

import (
	"fmt"
)
//类型断言
// 首先类型assertion和类型switch只能作用于interface{}，不能是其他类型
// 如果你输入错误它就继续让你输入这很明显就是个无限的for循环
func main(){
     // 声明一个变量，保存接收用户的输入选项
     key :=""
     loop :=true
     var balance float64 =10000.0 // 定义账户的余额
     var money float64 =0.0 // 每次的收支
     note :="" // 收支说明
     details :="\n收支\t\t\t账户余额\t\t\t收支\t\t说   明\t\n"
     // 默认初始化没有收支操作` 所以显示details无意义当有收支操作才进行显示
     // 涉及的编程思想:如果类行为发生了以后才出现A类行为 这时候我们得去加个标识符去判断或者加锁这种思想去判断
     // 锁就是与这类行为挂钩
     flag := false
     // 无限循环里面的程序直到退出
	 for {
		 fmt.Println("\n--------------家庭收支记帐软件--------------\n")
		 fmt.Println("1-收支明细")
		 fmt.Println("2-登记收入")
		 fmt.Println("3-登记支出")
		 fmt.Println("4-退出程序")
		 fmt.Println("请选择(1-4):")
		 // 注意传个指针过去 不然就改变不了key值 -- 很多时候我们要需要改变值 好奇怪的golang设计成值传递
		 _, _ = fmt.Scanln(&key)
		 switch key {
		 case "1":
			 fmt.Println("\n--------------当前收支明细记录--------------")
			 if flag{
				 fmt.Println(details)
			 }else{
				 fmt.Println("\n当前没有收支记录，来记一笔吧...\n")
			 }
		 case "2":
			 fmt.Println("本次收入金额：")
			 _, _ = fmt.Scanln(&money)
			 balance += money //修改账户余额
			 fmt.Println("本次收入说明：")
			 _, _ = fmt.Scanln(&note)
			 // 　%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名 %c 是以字符串输出 money是float类型
			 // Sprintf formats according to a format specifier and returns the resulting string. specifier说明符
			//  details += fmt.Sprintf("\n\n收入\t\t\t%c\t\t\t%v\t\t\t%v",balance,money,note) // %!c(float64=10100)
			 details += fmt.Sprintf("\n\n收入\t\t\t%v\t\t\t%v\t\t%v",balance,money,note)
			 flag=true
		 case "3":
			 fmt.Println("登记支出金额：")
			 fmt.Scanln(&money)
			 if money>balance{
			 	fmt.Println("余额不足,请重新输入....")
			 	break;
			 }
			 balance -= money
			 fmt.Println("本次支出说明：")
			 fmt.Scanln(&note)
			 details += fmt.Sprintf("\n\n支出\t\t\t%v\t\t\t%v\t\t%v",balance,money,note)
			 flag=true

		 case "4":
		 	fmt.Println("你确认要退出吗?y/n")
		 	 choice :=""
			 loop=false
			 //  如果没有输入n或者y就一直在这里显示所以需要个for无限循环
			 for{

                  fmt.Scanln(&choice)
                  if choice=="y"||choice=="n"{
                  	break
				  }
                  fmt.Println("您的输入有误，请重新输入y/n")
			 }
			 if(choice=="y"){
			 	loop=false
			 }

		 default:
		 	fmt.Println("请输入正确的选项...")

		 }
		 if !loop{
		 	break // 就会跳出循环
		 }

	 }
     fmt.Println("你退出了家庭记账软件的使用....")

}

// 加个功能不想输入支入支出