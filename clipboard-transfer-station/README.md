# clipboard-transfer-station/clipboard-transfer-station/README.md

# 项目愿景

## 简介
本项目旨在开发一款轻量级的基于服务器的文件/剪贴板中转站，允许用户在不同设备之间快速、安全地共享文件和剪贴板内容。

## 目标用户
本项目的目标用户包括需要在多设备间高效共享信息的个人用户和团队，尤其是开发者和设计师。

## 核心功能
- 跨设备文件和剪贴板内容共享
- 实时同步和状态更新
- 友好的用户界面和交互体验

# 技术栈

## 后端
- **语言**：Go
- **框架**：Gin（轻量级的Web框架，适合高性能需求）
- **数据存储**：SQLite（轻量级数据库，适合小型项目）

## Windows 客户端
- **框架**：Tauri（支持Fluent Design/Mica材质，轻量级）
- **前端**：React（使用TypeScript，适合构建现代UI）

## Android 客户端
- **框架**：Jetpack Compose（支持Material You，现代化UI开发）
- **语言**：Kotlin（Android开发的推荐语言）

# 模块划分

## 后端模块
- `main.go`：后端应用的入口文件，负责启动服务器并设置路由。
- `handlers/index.go`：定义处理请求的逻辑，包含API接口的实现。

## Windows 客户端模块
- `src-tauri/src/main.rs`：Tauri应用的主入口文件，负责初始化Tauri和设置窗口。
- `src/App.tsx`：React组件的主文件，负责渲染应用的UI。

## Android 客户端模块
- `app/src/main/java/com/transfer/MainActivity.kt`：Android应用的主活动文件，负责应用的启动和UI逻辑。

# 快速启动指南

## 后端启动步骤
1. 确保已安装Go环境。
2. 在`backend`目录下运行`go run src/main.go`启动后端服务。

## Windows 客户端启动步骤
1. 确保已安装Node.js和Rust环境。
2. 在`windows-client`目录下运行`npm install`安装依赖。
3. 在`windows-client/src-tauri`目录下运行`cargo tauri dev`启动Tauri应用。

## Android 客户端启动步骤
1. 确保已安装Android Studio和相关SDK。
2. 在`android-client`目录下打开项目，运行`MainActivity.kt`以启动Android应用。