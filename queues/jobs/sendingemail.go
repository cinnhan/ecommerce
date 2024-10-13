package jobs

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
	"myproject/constants"
)

func NewSendingEmail(email string) *asynq.Task {
	payload := map[string]interface{}{"email": email}

	return NewTask(constants.SendingEmailQueueName, payload)
}

func EnqueueSendingEmail(client *asynq.Client, email string) error {
	task := NewSendingEmail(email)

	if _, err := client.Enqueue(task, BuildTaskOption()); err != nil {
		return err
	}

	return nil
}

func HandleSendingEmail(ctx context.Context, task *asynq.Task) error {
	var mapPayload map[string]interface{}
	json.Unmarshal(task.Payload(), &mapPayload)
	email := mapPayload["email"]

	log.Printf("Sending email to %s\n", email)

	return nil
}
