dev:
	wails dev

build:
	wails build

linux:
	wails build -platform linux/arm64

windows:
	wails build -platform windows/amd64 -nsis

macos-amd64:
	wails build -platform darwin/amd64

macos-arm64:
	wails build -platform darwin/arm64

clean:
	rm -rf build/
	rm -rf frontend/pkg
	rm -rf frontend/src/lib/wailsjs
