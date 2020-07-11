package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

// Help Information.
var Usage = `bbhj 机器人使用帮助

		示例1：bbhj 4407
		示例2：bbhj 山西 4407
		示例3：bbhj 山西 运城 张

		说明：bbhj命令支持最多3个关键字的查询; 命令及各关键字只能以空格分隔。
`

// Refer: https://xuthus.cc/go/scf-go-runtime.html
func handler(ctx context.Context, event events.TimerEvent) {
	fmt.Println("hello", time.Now())
	return
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(handler)
}
