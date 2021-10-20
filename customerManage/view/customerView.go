package view

import (
	"customerManage/controller"
	"customerManage/model"
	"fmt"
)

// view层类

// CustomerView TODO
type CustomerView struct {
	key  string // 接收用户输入
	loop bool   // 表示是否循环显示菜单
	// controller控制器
	customerService *controller.CustomerService // 表示是controller.CustomerService类型的指针对象
}

// 实际不用声明构造函数也行 因为go直接导包实例化就行 但按照面向对象的习惯还是需要

// NewCustomerView TODO
// 声明构造函数
func NewCustomerView(key string, loop bool) *CustomerView {
	return &CustomerView{
		key,
		loop,
		controller.NewCustomerService(), // 这一块controll的实例对象
		// 这个实例化的结构体的方法接收者都是指针对象 也就是指向结构体的同一块内存 就并不是拷贝了
		// 因为接收者实际就是参数 go参数就是值传递（指针也是拷贝）---这个理解很重要！！！
	}
}

// MainMenu TODO
// 显示菜单
// golang的方法大小写决定了是否是私有方法 私有方法只能在类的内部使用 外部不可直接使用
func (this *CustomerView) MainMenu() {
	for {
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
			// fmt.Println("添 加 用 户")
			this.Add()
		case "2":
			// fmt.Println("修 改 用 户")
			this.Update()
		case "3":
			// fmt.Println("删 除 用 户")
			this.Delete()
		case "4":
			this.List()
		case "5":
			this.loop = false
		default:
			fmt.Println("您的输入有误,请重新输入....")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您退出了客户关系管理系统...\n")
}

// List TODO
// 拿到list去展示数据 这种写法代表结构体的方法 是可以拿到结构体里的属性的不管大小写 面向对象的内部方法当然都可以拿到
func (this *CustomerView) List() {
	// 获取list
	customers := this.customerService.SelectList()
	fmt.Println("---------------------------客户列表---------------------------\n")
	fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t\t电话\t\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		// 返回的是customers实例化结构体 肯定携带着自己的方法
		// 静态语言的编程思想很严重 就是限定死的 这种更好 比如你看到数组或者slice就与循环联想在一起
		fmt.Println(customers[i].GetInfo()) // 记得打印出来
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")

}

// 添加

// Add TODO
func (this *CustomerView) Add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewNotIdCustomer(name, gender, age, phone, email)
	// 调用control层添加数据
	// 又理解出来就是view就是做交互 交互层 比如我点击按钮添加数据 添加数据底层不是我做的 我负责调用就行 我只负责UI层让他调用

	if this.customerService.AddList(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

// 删除

// Delete TODO
func (this *CustomerView) Delete() {
	fmt.Println("\n----------------------删除客户-----------------------\n")
	fmt.Println("请选择待删除的客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id) // 输入字母赋值不了 因为是int类型
	if id == -1 {
		// fmt.Println("\n-----------------退出成功-----------------\n")
		return
	}
	choice := ""
	for {
		fmt.Println("是否确认删除(Y/N): ")

		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "N" || choice == "n" {
			break
		}
		fmt.Println("输入错误请重新输入")
	}
	if choice == "Y" || choice == "y" {
		this.customerService.DeleteList(id)
		fmt.Println("\n----------------------删除成功-----------------------\n")
	} else {
		fmt.Println("\n----------------------删除失败-----------------------\n")
	}

}

// 更新

// Update TODO
func (this *CustomerView) Update() {
	fmt.Println("请选择待更新的客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id) // 输入字母赋值不了 因为是int类型
	if id == -1 {
		// fmt.Println("\n-----------------退出成功-----------------\n")
		return
	}
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(id, name, gender, age, phone, email)
	this.customerService.UpdateList(id, customer)

}

// generic 通用

// 编程思想:对一列行为（比如异步处理）应该包含正确还有涉及错误的处理 特别是异步的处理 更需要

// 结构体方法接收者本身就是个参数 也就是值拷贝所以考虑指向同个内存就是要用指针类型接收者

// 读取文件不是一次性读取 而是拿一点放入缓冲区 等缓冲区没了再继续读取 写入也是一样 放到缓冲区再写入

// 内存泄漏用完内存没有释放 比如文件的打开和关闭

// 文件的flag很重要主要是以什么方式操作文件->返回的flie对象就具备什么样的权限

// bufio是一次一次读的 io是全部一次性读出来

// Scanln还是有bug的这个方法
