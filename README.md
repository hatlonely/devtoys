# README

DevToys is 是一个开发者工具集合. 其灵感来自于有名的 [DevToys](https://github.com/veler/DevToys) 项目。
使用 [Wails](https://wails.io/docs/introduction) 和 [Svelte](https://svelte.dev/) 开发，提供优美的界面和跨平台的能力。

## 安装

当前正在开发中，暂时还没有 release 版本，以下是通过源码安装的方式

1. 安装 [nodejs](https://nodejs.org/en/download)
2. 安装 [golang 开发环境](https://go.dev/dl/)
3. 安装 [wails 开发工具](https://wails.io/docs/gettingstarted/installation)

```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

4. 下载源代码并编译

```
git clone https://github.com/hatlonely/devtoys.git
make build
```

5. 打开 `buil/bin/devtoys` 运行程序

## 开发

参照安装步骤，安装必要的依赖后，执行 `wails dev` 即可进入开发模式。浏览器访问 `http://localhost:34115` 即可看到界面。

## 参考链接

- [DevToys](https://github.com/veler/DevToys)
- [Wails](https://wails.io/docs/introduction)
- [Svelte](https://svelte.dev/)
