package main

import (
	// "encoding/binary"
	"fmt"
	"log"

	"github.com/sstallion/go-hid"
)

func main() {

	// Initialize the hid package
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	d, err := hid.OpenFirst(0x0483, 0x5750)
	if err != nil {
		log.Fatal(err)
	}

	// Read the Manufacturer String.
	s, err := d.GetMfrStr()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Manufacturer String: %s\n", s)

	// Read the Product String.
	s, err = d.GetProductStr()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Product String: %s\n", s)

	// Read the Serial Number String.
	s, err = d.GetSerialNbr()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serial Number String: %s\n", s)

	const ENC_MAX = 200
	const ENC_DIV = 100
	var encoderValue float32 = 0

	for ;true; {
		buf := make([]byte, 6)
		_, err := d.Read(buf)
		if err != nil {
			log.Fatal(err)
		}	

		var button byte = buf[1]

		var diff int32
		diff |= int32(buf[5])
		diff |= int32(buf[4]) << 8
		diff |= int32(buf[3]) << 16
		diff |= int32(buf[2]) << 24

		encoderValue = encoderValue + (float32(diff) / ENC_DIV)
		if encoderValue <= 0 {
			encoderValue = 0
		} else if encoderValue >= ENC_MAX {
			encoderValue = ENC_MAX
		}

		var sendData uint32 = (uint32(button) << 16) | uint32(encoderValue)

		//fmt.Printf("DATA: [%d]%v\n", l, buf)
		fmt.Printf("\rENCODER: %05f\tBUTTON: %d\tOUT:%#x", encoderValue, button, sendData)
	}

	// err = d.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
