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
	customer :=model.NewCustomer(0,"张三","男",20,"15363383750","zs@souhu")
	customerService.customers=append(customerService.customers,customer)
	customer =model.NewCustomer(1,"王五","男",21,"15363383751","ww@souhu")
	customerService.customers=append(customerService.customers,customer)
	return customerService
}

// 查
func(this *CustomerService)SelectList()[]model.Customer{
	  return this.customers
}

// 方法不一定是为interface准备的面向对象的使用就是需要方法 interface是对继承的补充也是多态的体现

//  controller存放着数据 对数据进行操作


// 增 error存在没意义 就不返回了
func (this *CustomerService)AddList(customer model.Customer)bool{
	// fmt.Println(customers) //{0 vnues 男 22 15363383750 1518190293@qq.com}
	 this.customerNum++
	  customer.Id= this.customerNum
      this.customers=append(this.customers,customer)
      return true
}
// 删除
func (this *CustomerService)DeleteList(Id int)bool{

	   // 思路
	   // 根据结构体循环找出下标
	   index :=this.findById(Id)
	   // func append(slice []Type, elems ...Type) []Type 所以DeleteList也没必要返回error
	   this.customers=append(this.customers[:index],this.customers[index+1:]...)
	   return true

}



// 更新

func (this *CustomerService) UpdateList(Id int,customer model.Customer)bool{

	    this.customers[this.findById(Id)]=customer

	    return true
}


// 查找下标

func (this *CustomerService)findById(Id int) int{
	index :=-1
	for i,item:=range this.customers{
		if(Id==item.Id){
			index=i
			break
		}
	}
	return index
}