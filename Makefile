dev:
	wails dev

windows:
	wails build -platform windows/amd64 -nsis

macos-amd64:
	wails build -platform darwin/amd64

macos-arm64:
	wails build -platform darwin/arm64
