# Windows 主题切换器（精简版）

一个极简的 Windows 应用程序，一键切换系统主题颜色（深色/浅色模式）。

## 🎯 功能特性

- ✅ **一键切换**：对话框显示当前状态，点击按钮切换主题
- ✅ **极简界面**：最简单的Windows消息框，无需复杂界面
- ✅ **立即生效**：修改后系统立即更新，无需重启
- ✅ **绿色软件**：单个 .exe 文件，无需安装，无依赖
- ✅ **无管理员权限**：普通用户权限即可运行

## 📦 提供的版本

| 版本 | 文件 | 用途 | 大小 |
|------|------|------|------|
| **简单GUI版** | `bin\theme-simple.exe` | 推荐：双击运行，对话框显示状态并切换 | 2.5MB |
| **命令行版** | `bin\theme-cli.exe` | 开发者：命令行快速切换/查看状态 | 2.5MB |

## 🚀 立即使用

### 简单GUI版（推荐）
1. 打开 `theme-switcher\bin\` 文件夹
2. 双击 `theme-simple.exe`
3. 对话框中显示当前主题状态
4. 点击"是"切换到另一种主题
5. 点击"否"退出

**界面示例：**
```
┌─────────────────────┐
│       主题切换器     │
├─────────────────────┤
│ 当前主题: 浅色        │
│                     │
│ 点击确定切换到深色     │
│                     │
│     [是]  [否]       │
└─────────────────────┘
```

### 命令行版（快速）
```bash
cd theme-switcher\bin

# 查看当前主题
.\theme-cli.exe status

# 切换到浅色主题
.\theme-cli.exe light

# 切换到深色主题
.\theme-cli.exe dark

# 显示帮助
.\theme-cli.exe help
```

## 📁 项目结构

```
theme-switcher\
├── bin\                     # 可执行文件目录
│   ├── theme-simple.exe     # 简单GUI版本（推荐）
│   └── theme-cli.exe        # 命令行版本
├── scripts\                 # 构建脚本目录（可选）
├── simple-gui.go            # 简单GUI版本源代码
├── theme-cli.go             # 命令行版本源代码
├── go.mod                   # Go模块定义
├── go.sum                   # 依赖锁定文件
├── README.md                # 说明文档
└── LICENSE                  # MIT许可证
```

## 🔨 构建项目

### 使用构建脚本（最简单）
```bash
cd theme-switcher
# 运行构建菜单
build.bat
```

### 手动构建
```bash
cd theme-switcher
# 设置Go环境变量（一次性）
$env:GOROOT="D:\Program Files\Go"
$env:GOSUMDB="off"

# 编译简单GUI版本（隐藏控制台窗口）
go build -ldflags "-H windowsgui" -o bin\theme-simple.exe simple-gui.go

# 编译命令行版本
go build -o bin\theme-cli.exe theme-cli.go
```

## ⚙️ 技术原理

Windows 10/11 的主题设置存储在注册表中：
```
HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize

AppsUseLightTheme    = 1  // 应用使用浅色主题（0=深色）
SystemUsesLightTheme = 1  // 系统使用浅色主题（0=深色）
```

修改后发送 `WM_SETTINGCHANGE` 消息通知系统立即更新界面。

## ⚠️ 兼容性
- ✅ Windows 10 (1709+) 和 Windows 11
- ❌ Windows 7/8（不支持现代主题API）
- ✅ 普通用户权限（无需管理员）
- ✅ 64位和32位系统（Go自动适配）

## 📄 许可证
MIT License - 详见 [LICENSE](LICENSE) 文件。