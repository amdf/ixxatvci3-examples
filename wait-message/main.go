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

	//ID only
	msg, err = can.GetMsgByID(0x123, 5*time.Second)
	if nil == err {
		fmt.Printf("got msg ID: %03X\n", msg.ID)
	} else {
		fmt.Println(err.Error())
	}

	//ID and size
	msg, err = can.GetMsgByIDAndSize(0x123, 8, 5*time.Second)
	if nil == err {
		fmt.Printf("got ID: %03X, size: %d\n", msg.ID, msg.Len)
	} else {
		fmt.Println(err.Error())
	}

	//any message from list
	idlist := make(map[uint32]bool)
	idlist[0x123] = true
	idlist[0x456] = true
	idlist[0xABC] = true
	msg, err = can.GetMsgByIDList(idlist, 5*time.Second)
	if nil == err {
		fmt.Printf("got ID: %03X, size: %d\n", msg.ID, msg.Len)
	} else {
		fmt.Println(err.Error())
	}
}
