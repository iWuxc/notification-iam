package notification

import "github.com/marmotedu/iam/internal/apiserver/store"

type NotificationController struct {

}

func NewMsgPushController(store store.Factory) *NotificationController {
	return &NotificationController{

	}
}
