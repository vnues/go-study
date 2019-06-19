package view

import (
	"../controller"
	"fmt"
)

// view层类

type CustomerView struct {
   key string // 接收用户输入
   loop bool  // 表示是否循环显示菜单
   // controller控制器
   customerService *controller.CustomerService //表示是controller.CustomerService类型的指针对象
}
// 实际不用声明构造函数也行 因为go直接导包实例化就行 但按照面向对象的习惯还是需要

// 声明构造函数
func NewCustomerView(key string,loop bool)*CustomerView{
	 return &CustomerView{
	 	key,
	 	loop,
		 controller.NewCustomerService(),
	 }
}

// 显示菜单
// golang的方法大小写决定了是否是私有方法 私有方法只能在类的内部使用 外部不可直接使用
func (this *CustomerView) MainMenu(){
    for{
		fmt.Println("-----------------客户信息管理软件-----------------\n")
		fmt.Println("              1 添 加 客 户")
		fmt.Println("              2 修 改 客 户")
		fmt.Println("              3 删 除 客 户")
		fmt.Println("              4 客 户 列 表")
		fmt.Println("              5 退 出")
		fmt.Println("\n请选择(1-5): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添 加 用 户")
		case "2":
			fmt.Println("修 改 用 户")
		case "3":
			fmt.Println("删 除 用 户")
		case "4":
			this.List()
		case "5":
			this.loop=false
		default:
			fmt.Println("您的输入有误,请重新输入....")
		}
		if(!this.loop){
			break
		}
	}
    fmt.Println("您退出了客户关系管理系统...")
}

// 拿到list去展示数据 这种写法代表结构体的方法 是可以拿到结构体里的属性的不管大小写 面向对象的内部方法当然都可以拿到
func (this *CustomerView)List(){
	 // 获取list
	 customers :=this.customerService.SelectList()
	 fmt.Println("---------------------------客户列表---------------------------\n")
	 fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t\t电话\t\t\t邮箱")
	for i := 0; i < len(customers); i++ {
           // 返回的是customers实例化结构体 肯定携带着自己的方法
           // 静态语言的编程思想很严重 就是限定死的 这种更好 比如你看到数组或者slice就与循环联想在一起
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")

}