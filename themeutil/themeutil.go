package themeutil

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

func SetTheme(useLightTheme bool) error {
	key, err := registry.OpenKey(registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize`,
		registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("打开注册表失败: %w", err)
	}
	defer key.Close()

	value := uint32(0)
	if useLightTheme {
		value = 1
	}

	err = key.SetDWordValue("AppsUseLightTheme", value)
	if err != nil {
		return fmt.Errorf("设置 AppsUseLightTheme 失败: %w", err)
	}

	err = key.SetDWordValue("SystemUsesLightTheme", value)
	if err != nil {
		return fmt.Errorf("设置 SystemUsesLightTheme 失败: %w", err)
	}

	broadcastThemeChange()

	return nil
}

func GetCurrentTheme() (bool, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize`,
		registry.QUERY_VALUE)
	if err != nil {
		return false, fmt.Errorf("打开注册表失败: %w", err)
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return false, fmt.Errorf("读取主题值失败: %w", err)
	}

	return val == 1, nil
}

func GetThemeName(isLight bool) string {
	if isLight {
		return "浅色"
	}
	return "深色"
}

func broadcastThemeChange() {
	const (
		HWND_BROADCAST   = 0xFFFF
		WM_SETTINGCHANGE = 0x001A
		SMTO_ABORTIFHUNG = 0x0002
	)

	user32 := syscall.NewLazyDLL("user32.dll")
	sendMessageTimeout := user32.NewProc("SendMessageTimeoutW")

	policyStr := "Policy"
	ptr, _ := syscall.UTF16PtrFromString(policyStr)

	sendMessageTimeout.Call(
		uintptr(HWND_BROADCAST),
		uintptr(WM_SETTINGCHANGE),
		0,
		uintptr(unsafe.Pointer(ptr)),
		uintptr(SMTO_ABORTIFHUNG),
		1000,
		0,
	)
}
