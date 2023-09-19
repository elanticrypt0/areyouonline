package main

import (
	"fmt"
	"time"

	"github.com/elantycrypt0/go4it/src"
	"github.com/go-co-op/gocron"
)

func main() {
	fmt.Println("Buchón is running")

	server := src.NewServer("www.someweb.com", "www.somewebalternative.net")
	notify := src.NewNotify("emailme@youremail.com", "This message will be send")

	cron := gocron.NewScheduler(time.UTC)
	cron.StartAsync()
	job, err := cron.Every(30).Seconds().Do(func() {
		server.CheckServerStatus(notify)
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", job)
	}

	cron.StartBlocking()
}
