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
	can, err := b.Speed(ixxatvci3.Bitrate25kbps).Mode("29bit").Get()

	if err != nil {
		log.Fatal(err.Error())
	}

	can.Run()
	defer can.Stop()
	fmt.Println("Start")

	var msg1 candev.Message
	var msg2 candev.Message

	msg1.ID = 0x123
	msg1.Len = 3
	msg1.Ext = true //send short ID in 29bit mode
	msg1.Data[0] = 0x11
	msg1.Data[1] = 0x22
	msg1.Data[2] = 0x33

	msg2 = msg1
	msg2.ID = 0x55555 //29bit ID

	for {
		can.Send(msg1)
		time.Sleep(time.Second / 4)
		can.Send(msg2)
		time.Sleep(time.Second / 4)
	}

}
