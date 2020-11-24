package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	background := context.Background()

	fmt.Print(background.Deadline())

	ctx, cancelFunc := context.WithCancel(background)
	go dostuf(ctx)
	time.Sleep(20 * time.Second)
	cancelFunc()
	log.Print("Down")

	select {}
}

func dostuf(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Print("Done")
			return
		default:
			log.Print("WORKing....")
		}
	}
}
