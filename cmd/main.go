package main

import (
	"fmt"
	"github.com/shahid-io/orbit-queue/scheduler"
	"github.com/shahid-io/orbit-queue/api"
)

func main() {
	fmt.Println("Job Scheduler starting up...")
	scheduler.StartSchedular()
	api.StartServer()

	// TODO: Initialize router, scheduler, and other components
}
