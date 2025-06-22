package utils

import (
	"context"
	"log"
	"sync"
)

func SafeGo(ctx context.Context, wg *sync.WaitGroup, taskName string, task func(ctx context.Context)) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGo] Recovered from panic in task '%s': %v", taskName, r)
			}
		}()

		task(ctx)
	}()
}
