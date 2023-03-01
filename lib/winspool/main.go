package main

//#include "dllmain.h"
import (
	"C"
	"log"
	"os"
	"unsafe"

	"github.com/sstallion/go-hid"
)

const VID = 0x0483
const PID = 0x5750

const ENC_MAX = 200
const ENC_DIV = 100

const logFileName = "injection.log"

var hidDevice *hid.Device = nil
var logFile *os.File = nil
var encoderValue float32 = 0

//export OnProcessAttach
func OnProcessAttach() {
	var err error
	logFile, err = os.OpenFile(logFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logFile = nil
	}
	log.SetOutput(logFile)

	err = hid.Init()
	if err != nil {
		if logFile != nil {
			log.Printf("HID Init error: %s\n", err)
		}
		return
	}

	hidDevice, err = hid.OpenFirst(VID, PID)
	if err != nil {
		if logFile != nil {
			log.Printf("HID Open device %#x/%#x error: %s\n", VID, PID, err)
		}
		return
	}

}

//export ClosePrinter
func ClosePrinter() { }

//export OpenPrinterA
func OpenPrinterA() { }

// const MS_CTS_ON = 0x0010 // EAX | Enc 1
// const MS_DSR_ON = 0x0020 // r11d | BTN ?
// const MS_RING_ON = 0x0040 // r8d | Enc 2
// const MS_RLSD_ON = 0x0080

//export DocumentPropertiesA
func DocumentPropertiesA(handle unsafe.Pointer, staus *uint32) int {
	if hidDevice == nil {
		*staus = 0
		return 1
	}

	buf := make([]byte, 6)
	_, err := hidDevice.Read(buf)
	if err != nil {
		if logFile != nil {
			log.Printf("HID Device %v/%v readding error: %s\n", VID, PID, err)
		}
		*staus = 0
		return 1
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

	*staus = (uint32(button) << 16) | uint32(encoderValue)
	return 1
}

func main(){ }