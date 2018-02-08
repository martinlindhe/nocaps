package nocaps

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

var (
	keyPath = `SYSTEM\CurrentControlSet\Control\Keyboard Layout`
	keyName = `Scancode Map`
)

func EnableCaps() error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.WRITE)
	if err != nil {
		return err
	}
	err = key.DeleteValue(keyName)
	fmt.Println("Deleted registry at", keyPath+`\`+keyName)
	fmt.Println("You need to reboot for the settings to take effect")
	return err
}

func DisableCaps() error {
	newK, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	defer newK.Close()
	if err != nil {
		return err
	}

	err = newK.SetBinaryValue(keyName, []byte{
		0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x3a, 0x0,
		0x0, 0x0, 0x0, 0x0,
	})
	if err != nil {
		return err
	}

	fmt.Println("Updated registry at", keyPath+`\`+keyName)
	fmt.Println("You need to reboot for the settings to take effect")
	return nil
}
