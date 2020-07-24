package main

import (
	"../model"
	"../service"
	"fmt"
	"os"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView) list(){
	customers := this.customerService.List()
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}
func (this *customerView) add(){
	fmt.Println("---------------------添加客户---------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := int8(0)
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustimer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}
func (this *customerView)update()  {
	id := -1
	fmt.Scanln(&id)
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := int8(0)
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustimer(id,name, gender, age, phone, email)
	this.customerService.Update(customer)
}
func (this *customerView)delete()  {
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}else {
		this.customerService.Delete(id)
	}


}
func (this *customerView) manMenu(){
	for true {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			os.Exit(1)
		default:
			fmt.Println("Error input")




		}
		if !this.loop {
			break
		}
	}
	fmt.Println("exit(0)")
}
func main() {

	customerView := customerView{
		key:             "",
		loop:            true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.manMenu()

}