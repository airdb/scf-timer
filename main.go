package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Run)
}

// Refer: https://xuthus.cc/go/scf-go-runtime.html
func Run(ctx context.Context, event events.TimerEvent) {
	fmt.Println("hello", time.Now())
	return
}

