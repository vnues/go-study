package main

import (
	"customerManage/view"
)

// mvc方式 先写model 再写Controller控制器  再写视图 view
// view<->control->model 这样的流程关系
func main() {
	customerView := view.NewCustomerView("", true)
	// 在main函数中创建一个customerView并运行主菜单
	customerView.MainMenu()
}
