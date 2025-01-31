package main

import (
	days_server "github.com/billyarnold/outit-planner-be/days"
	trips_server "github.com/billyarnold/outit-planner-be/trips"
)

func main() {
	go days_server.Start()
	go trips_server.Start()

	select {}
}
