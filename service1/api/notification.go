package api

import (
	"context"
	"encoding/json"
	"net/http"
	"service1/proto/mail"
	"time"

	"go.opencensus.io/trace"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// NotificationHandler for SubscriptionNotification
func (api *API) NotificationHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//////////////////////////////////////////// Trace //////////////////////////////////////////////////////////////////////////
	ctx, span := trace.StartSpan(context.Background(), "NotificationHandler")
	span.AddAttributes(trace.StringAttribute("service1", "NotificationHandler"))
	logrus.Printf("TraceId: %s\n", span.SpanContext().TraceID.String())
	defer span.End()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	txMail := &mail.MailTransaction{
		Id:             span.SpanContext().TraceID.String(),
		SubscriptionId: "SubscriptionId",
		TxCreatedAt:    time.Now().Unix(),
		ReportedBy:     "System",
	}

	sub, err := api.sp.GetSubscription(ctx, txMail.SubscriptionId)
	if err != nil {
		logrus.Errorf("GetSubscription: error: %s", err)
		return
	}
	txMail.UserId = sub.UserId

	time.Sleep(1 * time.Second)

	spanContextJson, err := json.Marshal(span.SpanContext())
	if err != nil {
		return
	}
	txMail.Trace = string(spanContextJson)
	err = api.pub.Publish(txMail)
	if err != nil {
		logrus.Errorf("Publish: error: %s", err)
	}

	JSON(w, map[string]interface{}{
		"status": "ok",
	})
}
