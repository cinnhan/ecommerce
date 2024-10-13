package main

import (
	"fmt"
	"myproject/routes"
	"net/http"
	"os"
)

func main() {
	router := routes.BuildApiRouter()

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
