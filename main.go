package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	path := `SYSTEM\CurrentControlSet\Control\Keyboard Layout`
	keyName := `Scancode Map`
	newK, _, err := registry.CreateKey(registry.LOCAL_MACHINE, path, registry.SET_VALUE)
	defer newK.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = newK.SetBinaryValue(keyName, []byte{
		0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x3a, 0x0,
		0x0, 0x0, 0x0, 0x0,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated registry at", path+`\`+keyName)
	fmt.Println("You need to reboot for the settings to take effect")
}
