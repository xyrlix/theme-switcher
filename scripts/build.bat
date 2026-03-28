@echo off
echo ========================================
echo  Windows Theme Switcher - Build Tool
echo ========================================
echo.

REM Set Go environment variables
set "GOROOT=D:\Program Files\Go"
set "GOSUMDB=off"

REM Get current directory
set "SCRIPT_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPT_DIR%.."

REM Create output directory
if not exist "%PROJECT_DIR%\bin" mkdir "%PROJECT_DIR%\bin"

REM Change to project root directory
cd /d "%PROJECT_DIR%"
echo Current directory: %CD%
echo.

REM Run full build directly
goto full

:full
echo Full Build Mode...
echo.

REM 1. Clean old builds
echo 1. Cleaning old builds...
del bin\*.exe 2>nul
echo OK: Clean completed

REM 2. Download dependencies
echo.
echo 2. Downloading dependencies...
go mod download
if %errorlevel% neq 0 (
    echo ERROR: Dependency download failed
    exit /b 1
)
echo OK: Dependencies downloaded

REM 3. Build CLI version
echo.
echo 3. Building CLI version...
go build -o bin\theme-cli.exe theme-cli.go
if %errorlevel% neq 0 (
    echo ERROR: CLI build failed
    exit /b 1
)
echo OK: theme-cli.exe built

REM 4. Build simple GUI version
echo.
echo 4. Building simple GUI version...
go build -ldflags "-H windowsgui" -o bin\theme-switcher.exe simple-gui.go
if %errorlevel% neq 0 (
    echo ERROR: Simple GUI build failed
    set "GUI_ERROR=1"
) else (
    echo OK: theme-switcher.exe built
)

echo.
echo ========================================
echo Full build completed!
echo.  
echo Available versions:
echo   - bin\theme-cli.exe      (CLI version)
if not defined GUI_ERROR (
    echo   - bin\theme-switcher.exe  (Simple GUI version)
)
echo ========================================
goto :end

:end
pause
