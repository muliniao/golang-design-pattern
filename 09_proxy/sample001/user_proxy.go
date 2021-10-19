package sample001

import (
	"fmt"
	"time"
)

type UserControllerProxy struct {
	userController		*UserController
	metricsController	*MetricsController
}

func NewUserControllerProxy(userController *UserController) *UserControllerProxy {
	return &UserControllerProxy{
		userController:    userController,
		metricsController: &MetricsController{},
	}
}

func (u *UserControllerProxy) Login(telephone string, password string) {
	fmt.Println("userControllerProxy login is calling")

	before := time.Now().UTC().Unix()
	u.userController.Login(telephone, password)
	after := time.Now().UTC().Unix()

	delta := after - before
	u.metricsController.recordRequest(string(delta))
}

func (u *UserControllerProxy) Register(telephone string, password string) {
	fmt.Println("userControllerProxy register is calling")

	before := time.Now().UTC().Unix()
	u.userController.Register(telephone, password)
	after := time.Now().UTC().Unix()

	delta := after - before
	u.metricsController.recordRequest(string(delta))
}
