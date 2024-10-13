package commands

import (
	"log"
	"myproject/queues/jobs"
)

func sampleCommand() {
	client := jobs.InitializeQueueClient()

	weekName := "Week 5"
	log.Printf("Getting report from week: %s\n", weekName)
	jobs.EnqueueWeeklyReport(client, weekName)
}
