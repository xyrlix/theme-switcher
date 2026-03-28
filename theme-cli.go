package main

import (
	"fmt"
	"os"
	"strings"

	"theme-switcher/themeutil"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Windows 主题切换器 - 命令行版")
		fmt.Println("使用方法:")
		fmt.Println("  theme-cli.exe light   # 切换到浅色主题")
		fmt.Println("  theme-cli.exe dark    # 切换到深色主题")
		fmt.Println("  theme-cli.exe status  # 查看当前主题状态")
		fmt.Println("\n示例: theme-cli.exe light")
		os.Exit(1)
	}

	cmd := strings.ToLower(os.Args[1])

	switch cmd {
	case "light":
		err := themeutil.SetTheme(true)
		if err != nil {
			fmt.Printf("切换到浅色主题失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ 已切换到浅色主题")

	case "dark":
		err := themeutil.SetTheme(false)
		if err != nil {
			fmt.Printf("切换到深色主题失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ 已切换到深色主题")

	case "status":
		isLight, err := themeutil.GetCurrentTheme()
		if err != nil {
			fmt.Printf("获取主题状态失败: %v\n", err)
			os.Exit(1)
		}
		if isLight {
			fmt.Println("📱 当前主题: 浅色 (Light)")
		} else {
			fmt.Println("🌙 当前主题: 深色 (Dark)")
		}

	default:
		fmt.Printf("未知命令: %s\n", cmd)
		fmt.Println("请使用: light, dark, 或 status")
		os.Exit(1)
	}
}
