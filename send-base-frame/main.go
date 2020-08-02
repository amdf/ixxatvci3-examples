package main

import (
	"fmt"
	"log"
	"time"

	"github.com/amdf/ixxatvci3"
	"github.com/amdf/ixxatvci3/candev"
)

func main() {
	var b candev.Builder
	can, err := b.Speed(ixxatvci3.Bitrate25kbps).Mode("11bit").Get()

	if err != nil {
		log.Fatal(err.Error())
	}

	can.Run()
	defer can.Stop()
	fmt.Println("Start")

	var msg candev.Message

	msg.ID = 0x123
	msg.Len = 3
	msg.Data[0] = 0x11
	msg.Data[1] = 0x22
	msg.Data[2] = 0x33

	for {
		can.Send(msg)
		time.Sleep(time.Second / 4)
	}

}
