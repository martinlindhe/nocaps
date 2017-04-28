package nocaps

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func Perform() error {
	path := `SYSTEM\CurrentControlSet\Control\Keyboard Layout`
	keyName := `Scancode Map`
	newK, _, err := registry.CreateKey(registry.LOCAL_MACHINE, path, registry.SET_VALUE)
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

	fmt.Println("Updated registry at", path+`\`+keyName)
	fmt.Println("You need to reboot for the settings to take effect")
	return nil
}
