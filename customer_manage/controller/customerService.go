package controller

import (
	"../model"
)
// model层类 对象肯定包括方法的 别太对mvc的片面理解 认为Model层只有属性没有方法
type CustomerService struct {
	customers [] model.Customer
	// 声明一个字段，表示当前切片含有多少用户
	customerNum int
}


// 构造函数
func NewCustomerService()*CustomerService{
    customerService :=&CustomerService{}
    // 问题一 如果不给customerService属性赋值会怎么样 应该都是初始化的空值
	customerService.customerNum=1 // 写法可以这样写
	// 注意传递函数参数和结构体属性实例化的区别 参数是必须的不然会报错，结构体的是属性
	// 默认初始化的值 如果有数据库会从数据库拿到
	customer :=model.NewCustomer(1,"张三","男",20,"15363383750","zs@souhu")
	customerService.customers=append(customerService.customers,customer)
	customer =model.NewCustomer(2,"王五","男",21,"15363383751","ww@souhu")
	customerService.customers=append(customerService.customers,customer)
	return customerService
}

// 查
func(this *CustomerService)SelectList()[]model.Customer{
	  return this.customers
}


// 方法不一定是为interface准备的面向对象的使用就是需要方法 interface是对继承的补充也是多态的体现