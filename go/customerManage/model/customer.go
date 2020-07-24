package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int8
	Phone  string
	Email  string
}

func NewCustimer(id int,name string,gender string,age int8,phone string,
email string) Customer{
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustimer2(name string,gender string,age int8,phone string,
	email string) Customer{
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}


func (this *Customer)GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id,
		this.Name, this.Gender,this.Age, this.Phone, this.Email)
	return info

}
