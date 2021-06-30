package main

func main() {
	adbinit(1)
	haveOneImgsClick(1, 0.05, false, getSystemImg("test.png"))
	// AdbShellScreencapPullRm()
	// rigorous_bitmap := robotgo.OpenBitmap("screen.png")
	// defer robotgo.FreeBitmap(rigorous_bitmap)
	// dst_map := robotgo.OpenBitmap(getSystemImg("test.png"))
	// defer robotgo.FreeBitmap(dst_map)
	// fx, fy := robotgo.FindBitmap(dst_map, rigorous_bitmap, 0.01)
	// fmt.Println(fx)
	// fmt.Println(fy)
	for false {
		//初始化ID
		infoID()
		//輸入想要處理的事情
	}
}
