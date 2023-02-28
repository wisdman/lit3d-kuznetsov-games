package main

import (
	"encoding/binary"
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

	for ;true; {
		buf := make([]byte, 5)
		l, err := d.Read(buf)
		if err != nil {
			log.Fatal(err)
		}	


		value := binary.BigEndian.Uint32(buf[1:]);
		fmt.Printf("\rVALUE: %+v DATA: [%d]%v   ", value, l, buf)

		// value := binary.BigEndian.Uint32(buf[1:4])


		// cnt := int16(binary.BigEndian.Uint16(buf[1:4]))
		// fmt.Printf("\rBTN:%d VALUE: %+v DATA: [%d]%v           ", buf[1], value, l, buf);
	}


	// buf := make([]byte, 5)
	// l, err := d.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Repport [%d]%v\n", l, buf);


	// err = d.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}