package utils

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestSafeGo_Success(t *testing.T) {
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	executed := false

	SafeGo(ctx, wg, "test-task", func(ctx context.Context) {
		executed = true
	})

	wg.Wait()

	if !executed {
		t.Error("Expected task to be executed")
	}
}

func TestSafeGo_WithPanic(t *testing.T) {
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	// Этот тест проверяет, что паника не приведет к краху программы
	SafeGo(ctx, wg, "panic-task", func(ctx context.Context) {
		panic("test panic")
	})

	// Если мы дошли до этой точки, значит паника была обработана
	wg.Wait()

	// Тест успешен, если мы дошли до этой точки без краха
}

func TestSafeGo_WithContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	taskStarted := false

	SafeGo(ctx, wg, "cancellation-task", func(ctx context.Context) {
		taskStarted = true
		select {
		case <-ctx.Done():
			// Контекст был отменен
			return
		case <-time.After(100 * time.Millisecond):
			// Таймаут
			return
		}
	})

	// Отменяем контекст
	cancel()
	wg.Wait()

	if !taskStarted {
		t.Error("Expected task to be started")
	}
}

func TestSafeGo_MultipleGoroutines(t *testing.T) {
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	counter := 0
	mu := &sync.Mutex{}

	// Запускаем несколько горутин
	for i := 0; i < 5; i++ {
		SafeGo(ctx, wg, "counter-task", func(ctx context.Context) {
			mu.Lock()
			counter++
			mu.Unlock()
		})
	}

	wg.Wait()

	if counter != 5 {
		t.Errorf("Expected counter to be 5, got %d", counter)
	}
}
