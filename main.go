package main

import "github.com/go-vgo/robotgo"

var status = 0
var notthink = 0 //多少次迴圈沒動作 1000次沒動作 就關閉重啟
var configObj ExampleConfig

func main() {
	infoID()
	adbinit(1)
	tmpconfig, _ := LoadSmallRoleConfig()
	tmpconfig.Name = "" //故意改變讓一開始進入回主選單
	for {
		//如果config有變動 需要重新回到主頁
		configObj, _ = LoadSmallRoleConfig()
		if tmpconfig != configObj {
			tmpconfig = configObj
			adbinit(configObj.Id)
			status = 0
		}

		//重開遊戲
		if status == 0 {
			//關閉遊戲
			closeApp()
			robotgo.Sleep(2)
			//啟動遊戲
			openApp()

			status = 1
		}

		//開啟後事情
		if status == 1 {
			toMainImg := []string{}
			toMainFunc := []func(x int, y int){}
			toMainImg, toMainFunc = addEx(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addExFunc(toMainImg, toMainFunc)
			haveOneImgsExecFunc(1, 0.05, false, toMainImg, toMainFunc...)
		}

		if notthink > 1000 {
			savescreen("notthink")
			status = 0
			notthink = 0
		}
		notthink++
	}
}
