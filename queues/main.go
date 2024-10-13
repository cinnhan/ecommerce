package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"myproject/constants"
	"myproject/queues/jobs"
)

func main() {
	client := jobs.InitializeQueueClient()

	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("Error closing redis connection: %v\n", err)
		}
	}()

	server := jobs.InitializeQueueServer()

	mux := asynq.NewServeMux()

	// worker list
	mux.HandleFunc(constants.SendingEmailQueueName, jobs.HandleSendingEmail)
	mux.HandleFunc(constants.WeeklyReportQueueName, jobs.HandleWeeklyReport)

	if err := server.Run(mux); err != nil {
		fmt.Printf("Error starting redis server: %v\n", err)
	}

}
