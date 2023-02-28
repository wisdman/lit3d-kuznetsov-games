package main

//#include "dllmain.h"
import (
	"C"
	"encoding/binary"
	"fmt"
	"os/user"
	"runtime"
	"unsafe"

	"github.com/sstallion/go-hid"
	"golang.org/x/sys/windows"
)
import (
	"io/ioutil"
)

const template = `
FnCall: %s
WorkDir: %s
CmdLine: %s
Arch: %s
User: %s
Integrity: %s
`

var d *hid.Device

//export OnProcessAttach
func OnProcessAttach(
	hinstDLL unsafe.Pointer, // handle to DLL module
	fdwReason uint32, // reason for calling function
	lpReserved unsafe.Pointer, // reserved
) {
	err := hid.Init()
	if err != nil {
		message := fmt.Sprintf("Error: %s", err);
		_ = ioutil.WriteFile("error.txt", []byte(message), 0644)
	}
	d, err = hid.OpenFirst(0x0483, 0x5750)
	if err != nil {
		message := fmt.Sprintf("Error: %s", err);
		_ = ioutil.WriteFile("error.txt", []byte(message), 0644)
	}

	alert()
}

func alert() {
	imageName, path, cmdLine := hostingImageInfo()
	title := fmt.Sprintf("Host Image: %s", imageName)
	arch := runtime.GOARCH

	usr, err := user.Current()
	if err != nil {
		usr.Username = "Unknown Error"
	}

	integrity, err := getProcessIntegrityLevel()
	if err != nil {
		integrity = "Unknown Error"
	}

	msg := fmt.Sprintf(template, caller(), path, cmdLine, arch, usr.Username, integrity)
	MessageBox(title, msg, MB_OK|MB_ICONEXCLAMATION|MB_TOPMOST)
}

func hostingImageInfo() (imageName, path, cmdLine string) {
	peb := windows.RtlGetCurrentPeb()
	userProcParams := peb.ProcessParameters
	imageName = userProcParams.ImagePathName.String()
	path = userProcParams.CurrentDirectory.DosPath.String()
	cmdLine = userProcParams.CommandLine.String()
	return
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
	buf := make([]byte, 5)
	_, err := d.Read(buf)
	if err != nil {
		message := fmt.Sprintf("Error: %s", err);
		_ = ioutil.WriteFile("error.txt", []byte(message), 0644)
		*staus = 0
		return 1
	} 

	*staus = binary.BigEndian.Uint32(buf[1:]);

	// if buf[1] == 1 {
	// 	*staus = MS_DSR_ON
	// } else {
	// 	*staus = 0
	// }

	return 1
}

func main(){ }