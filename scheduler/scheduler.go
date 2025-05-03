package scheduler

import (
	"github.com/robfig/cron/v3"
	"log"
	"sync"
	"fmt"
)

type JobRequest struct {
	ID string `json:"id"`
	Schedule string `json:"schedule"`
	Command string `json:"command"`
}


var (
	c *cron.Cron
	jobMutex sync.Mutex
	jobMap = make(map[string]cron.EntryID)
)



func StartSchedular(){
	c = cron.New()
	log.Println("Starting cron scheduler...")
	c.Start()
}

func RegisterJob(req JobRequest) error {
	jobMutex.Lock()
	defer jobMutex.Unlock()

	entryID, err := c.AddFunc(req.Schedule, func(){
		fmt.Println("[Job ID: %s] Executing command: %s\n", req.ID, req.Command)
	})
	if err != nil {
		return err
	} 

	jobMap[req.ID] = entryID
	log.Printf("Registered Job %s with Schedule %s\n", req.ID, req.Schedule)
	return nil
}


func StopScheduler() {
	c.Stop()
}