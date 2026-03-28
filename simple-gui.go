package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"theme-switcher/themeutil"
)

const (
	MB_OK       = 0x00000000
	MB_OKCANCEL = 0x00000001
	IDOK        = 1
	IDCANCEL    = 2
)

func main() {
	isLight, err := themeutil.GetCurrentTheme()
	if err != nil {
		msgBox("错误", fmt.Sprintf("获取当前主题失败: %v\n请手动运行。", err))
		os.Exit(1)
	}

	title := "🌙 主题切换器"
	message := fmt.Sprintf("当前主题: %s\n\n点击确定切换到 %s",
		themeutil.GetThemeName(isLight),
		themeutil.GetThemeName(!isLight))

	// 使用确定/取消按钮的消息框
	result := msgBoxOKCancel(title, message)

	// 只有点击确定时才切换主题
	if result == IDOK {
		err = themeutil.SetTheme(!isLight)
		if err != nil {
			msgBox("错误", fmt.Sprintf("切换主题失败: %v", err))
		}
	}
}

func msgBoxOKCancel(title, message string) int {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	ret, _, _ := messageBox.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(MB_OKCANCEL),
	)

	return int(ret)
}

func msgBox(title, message string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	messageBox.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(0x00000010), // MB_ICONERROR
	)
}
