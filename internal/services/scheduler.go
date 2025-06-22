package services

import (
	"context"
	"haircompany-shop-rest/pkg/utils"
	"log"
	"sync"
	"time"
)

type task struct {
	name     string
	callback func()
}

type Scheduler interface {
	CreateTask(name string, callback func()) *task
	StartEveryDay(hour, minute int, t *task)
}

type scheduler struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

func NewScheduler(ctx context.Context, wg *sync.WaitGroup) Scheduler {
	return &scheduler{
		ctx: ctx,
		wg:  wg,
	}
}

func (s *scheduler) CreateTask(name string, callback func()) *task {
	if name == "" {
		log.Printf("[Scheduler] Task name cannot be empty")
		return nil
	}
	if callback == nil {
		log.Printf("[Scheduler] Task callback cannot be nil for task: %s", name)
		return nil
	}

	return &task{
		name:     name,
		callback: callback,
	}
}

func (s *scheduler) StartEveryDay(hour, minute int, t *task) {
	utils.SafeGo(s.ctx, s.wg, t.name, func(ctx context.Context) {
		for {
			now := time.Now()
			nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

			if now.After(nextRun) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			wait := nextRun.Sub(now)
			log.Printf("[Scheduler] Task %s will run in %s", t.name, nextRun.Format("15:04:05"))

			select {
			case <-ctx.Done():
				log.Printf("[Scheduler] Stopping task %s due to context cancellation", t.name)
				return
			case <-time.After(wait):
				log.Printf("[Scheduler] Running task: %s", t.name)
				t.callback()
			}
		}
	})
}
