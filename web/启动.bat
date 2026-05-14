@echo off
chcp 65001 >nul
echo ========================================
echo   JHB积分商城 - 前端启动
echo ========================================
echo.
echo 正在打开首页...
echo.

start "" "%~dp0index.html"

echo ✓ 首页已打开
echo.
echo 提示：请确保后端服务已在 http://localhost:8888 运行
echo.
pause
