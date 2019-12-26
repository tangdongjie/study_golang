package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/royeo/dingrobot" //go get -u github.com/royeo/dingrobot
	"log"
	"unsafe"
)

const (
	a = "a"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main()  {
	c := cron.New()
	c.AddFunc("*/2 * * * * *",Iota)
	c.Start()
	Iota()

}

func Test()(int,int,string)  {
	a , b , c := 1 , 2 , "str"
	return a,b,c
}

func Iota()  {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	log.Println(a,b,c,d,e,f,g,h,i)

}



func SendDingMsg(msg string) {
	//请求地址模板
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=228f9afba4bedff3d62c0ebe214eaba929f4136651e33ada21a2f9b989c0b543"
	robot := dingrobot.NewRobot(webHook)
	atMobiles := []string{"18511626110","17600610625"}
	err := robot.SendText(msg,atMobiles,false)
	if err != nil {
		fmt.Println(err)
	}

}

