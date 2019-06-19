package model

import "fmt"

// 声明一个Customer结构体,表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 按照面向对象的模式 我们声明一个类 就需要有对应的构造函数 很明显go这里 我们需要自己自定义一个构造函数

// 使用一个工厂模式，返回一个Customer 声明一个构造函数
// 返回实例化的对象 (而且类定义的时候这样表示是这些属性必须传?)
func NewCustomer(id int,name string,gender string,age int,phone string,email string) Customer{
	     return Customer{
			 Id:id,
	     	 Name:name,
			 Gender:gender,
	     	 Age :age,
	     	 Phone:phone,
			 Email:email,
		 }
}

// model层提供返回用户信息的方法（view层筛选过滤做也行的 不用在意哪个哪个处理甚至在给control层做也行 但是control负责业务逻辑尽量保持简洁）
func(this *Customer)GetInfo() string{
    info :=fmt.Sprintf("\n%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
    return info
}