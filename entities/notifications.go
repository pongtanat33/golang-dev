package entities

type NotificationUseCase interface {
	GetUserNotifications(*UserNotificationsReq) (UserNotificationsRes, error)
	GetNotifications(*NotificationReq) (NotificationRes, error)
	CreateUserNotification(*UserNotificationsReq) error
}

type NotificationRepository interface {
	GetUserNotifications(*UserNotificationsReq) (UserNotificationsRes, error)
	GetNotifications(*NotificationReq) (NotificationRes, error)
	CreateUserNotification(*UserNotificationsReq) error
}

type UserNotificationsRes struct {
	Id              string `json:"id" db:"id"`
	UserEmail       string `json:"user_email" db:"user_email"`
	NotificationKey string `json:"notification_key" db:"notification_key"`
	NotificationId  string `json:"notification_id" db:"notification_id"`
	Seen            bool   `json:"seen" db:"seen"`
	SeenAt          string `json:"seen_at" db:"seen_at"`
}

type UserNotificationsReq struct {
	UserEmail      string `json:"user_email" db:"user_email"`
	NotificationId string `json:"notification_id" db:"notification_id"`
	Seen           bool   `json:"seen" db:"seen"`
}

type NotificationReq struct {
	NotificationKey string `json:"notification_key" db:"notification_key"`
	Type            string `json:"type" db:"type"`
	Version         string `json:"version" db:"version"`
}

type NotificationRes struct {
	NotificationKey string `json:"notification_key" db:"notification_key"`
	NotificationId  string `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	Message         string `json:"message" db:"message"`
	Description     string `json:"description" db:"description"`
	Type            string `json:"type" db:"type"`
	Version         string `json:"version" db:"version"`
}
