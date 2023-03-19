package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	controlChan := make(chan bool)
	go control(ctx, controlChan)
	// start the work function
	controlChan <- true
	controlChan <- true
	controlChan <- true
	time.Sleep(5 * time.Second)
	controlChan <- false
	fmt.Println("sleeping for 5 seconds - no work happening")
	time.Sleep(5 * time.Second)
	controlChan <- true
	time.Sleep(5 * time.Second)
	controlChan <- false
	fmt.Println("sleeping for 5 seconds - no work happening")
	time.Sleep(5 * time.Second)
	return nil
}

// controller will run in a goroutine and block.
// It will return when the context is cancelled.
// if a desired state is sent on the stateCtrl channel, it will either
// start or stop the workWork function.
func control(ctx context.Context, controlChan chan bool) {
	var (
		workCtx context.Context // context used to control work function
		cancel  context.CancelFunc
		running bool // flag to indicate whether work function is running
	)
	for {
		select {
		case <-ctx.Done():
			// context is cancelled, so stop the work function if it is running
			if cancel != nil {
				cancel()
			}
			return
		case start := <-controlChan:
			if start && !running {
				// start the work function
				workCtx, cancel = context.WithCancel(context.Background())
				go work(workCtx)
				running = true
			} else if !start && running {
				// stop the work function if it is running
				cancel()
				running = false
			}
		}
	}
}
func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[context cancelled]")
			return
		default:
			fmt.Println("[working]")
			// do some work
			time.Sleep(time.Second)
		}
	}
}
