package sample001

import "fmt"

type IUserController interface {
	Login(telephone string, password string)
	Register(telephone string, password string)
}

type UserController struct {

}

func (u *UserController) Login(telephone string, password string) {
	fmt.Println("userController login is calling")
}

func (u *UserController) Register(telephone string, password string) {
	fmt.Println("userController register is calling")
}

