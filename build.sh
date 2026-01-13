#!/usr/bin/env bash
set -e

APP_NAME="coder"
DIST_DIR="dist"

echo "🔨 Building $APP_NAME ..."

# 创建输出目录
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

# 强制关闭 cgo，保证可跨平台运行
export CGO_ENABLED=0

echo "▶ macOS (Intel x86_64)"
GOOS=darwin GOARCH=amd64 go build -o "$DIST_DIR/${APP_NAME}-mac-intel"

echo "▶ macOS (Apple Silicon arm64)"
GOOS=darwin GOARCH=arm64 go build -o "$DIST_DIR/${APP_NAME}-mac-arm"

echo "▶ Windows (x86_64)"
GOOS=windows GOARCH=amd64 go build -o "$DIST_DIR/${APP_NAME}-win.exe"

echo ""
echo "✅ Build finished. Files:"
ls -lh "$DIST_DIR"
