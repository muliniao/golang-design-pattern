package sample

type NotificationObserver struct {
	notificationService NotificationService
}

func (n * NotificationObserver) HandleRegisterSuccess(userID int) {
	n.notificationService.sendInboxMessage(userID, "welcome...")
}




type NotificationService struct {

}

func (n *NotificationService) sendInboxMessage(userID int, message string) {

}

