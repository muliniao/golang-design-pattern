package sample

type UserController struct {
	userService UserService
	list []RegObserver
}

func (u *UserController) register(telephone string, password string) int {

	userID := u.userService.register(telephone, password)

	for _, v := range u.list {
		v.HandleRegisterSuccess(userID)
	}

	return userID
}

func (u *UserController) addObserver(observer RegObserver) {
	u.list = append(u.list, observer)
}




type UserService struct {

}

func (u *UserService) register(telephone string, password string) int {
	return 12345
}
