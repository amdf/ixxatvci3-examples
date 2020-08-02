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
	msg.Rtr = true //request
	for {
		can.Send(msg)
		time.Sleep(time.Second / 4)
	}

}
