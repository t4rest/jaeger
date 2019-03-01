package api

import (
	"net/http"
	"service1/proto/mail"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// NotificationHandler for SubscriptionNotification
func (api *API) NotificationHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	txMail := &mail.MailTransaction{
		Id:             "Id",
		SubscriptionId: "SubscriptionId",
		TxCreatedAt:    time.Now().Unix(),
		ReportedBy:     "System",
	}

	sub, err := api.sp.GetSubscription(txMail.SubscriptionId)
	if err != nil {
		logrus.Errorf("GetSubscription: error: %s", err)
		return
	}

	txMail.UserId = sub.UserId

	err = api.pub.Publish(txMail)
	if err != nil {
		logrus.Errorf("Publish: error: %s", err)
	}

	JSON(w, map[string]interface{}{
		"status": "ok",
	})
}
