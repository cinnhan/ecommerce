package main

import (
	"myproject/console/commands"
)

func main() {
	// initializes and starts cron
	commands.InitializeCron()

	// worker keeps running to allow crontab to execute
	select {}
}
