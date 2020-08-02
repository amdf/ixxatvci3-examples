package main

import (
	"fmt"
	"log"

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

	ch, idx := can.GetMsgChannelCopy()
	defer can.CloseMsgChannelCopy(idx)

	for msg := range ch {
		fmt.Printf("ID: %03X, size: %d, rtr: %t, data: ", msg.ID, msg.Len, msg.Rtr)
		if msg.Len > 0 {
			for i := uint8(0); i < msg.Len; i++ {
				fmt.Printf(" %02X", msg.Data[i])
			}
		}
		fmt.Println()
	}
}
