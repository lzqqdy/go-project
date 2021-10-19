package crontab

import (
	"go-project/pkg/timer"
)

// Timer 启动定时器
func Timer() {
	crontab := timer.NewCrontab()

	//if err := crontab.AddByFunc("1", "*/1 * * * * ?", job.Test); err != nil {
	//	fmt.Printf("error to add crontab task:%s", err)
	//	os.Exit(-1)
	//}
	//.....
	crontab.Start()
	//select {}
}
