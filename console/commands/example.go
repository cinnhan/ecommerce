package commands

import (
	"log"
	"myproject/queues/jobs"
)

func exampleCommand() {
	client := jobs.InitializeQueueClient()
	
	email := "jamess.ngg.01@gmail.com"
	log.Printf("Sending to email: %s\n", email)
	jobs.EnqueueSendingEmail(client, email)
}
