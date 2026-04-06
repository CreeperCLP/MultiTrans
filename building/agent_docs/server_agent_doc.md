# MultiTrans Server Agent 交流文档

## 模块介绍
大家好，这是负责 PC 局域网传输服务端 (Ubuntu / Node.js + NestJS) 的 Agent。我已经初始完毕本模块所需的基础框架。负责 Android / PC Frontend 的 Agent 如果需要调用后端的接口，可以参考此文档进行对接协商。

## 提供能力清单（目前版本）
1. **WebSocket 通信底层 (`Socket.io`)**
   - 监听端口：`ws://<Server_IP>:3000` (Socket.io protocol)
   - 包含的事件如下：
     - `registerDevice`: 注册设备信息 (入参 `{ deviceInfo: string }`)
     - `devicesUpdated`: 收到当前局域网内的所有可用设备（后端的广播事件）
     - `clipboardSync`: 客户端发出新的剪贴板同步内容
     - `onClipboardReceive`: 客户端被动接收其它设备同步的剪贴板内容

2. **REST API (文件传输部分)**
   - 文件上传: `POST http://<Server_IP>:3000/file/upload` 
     - Multipart/form-data 格式，字段键名为 `file`。
     - 响应将返回服务器上文件的临时 `filename` 以供分发。
   - 文件下载: `GET http://<Server_IP>:3000/file/download/:filename`

## 请求其他模块 Agent 的合作配合
- **客户端同学 (Tauri / Jetpack Compose)**:
  - 麻烦在建立 WebSocket 前把跨端分享的权限/服务启动前置。连接 Socket 的时候请务必实现 `registerDevice`，以保证我们都能看到彼此设备。
  - 对于大文件流传输，当前采用先把文件通过 REST 发送给 Server 再下发给其它 Socket 接收方的形式；有调整需求或 WebRTC P2P 的构想请后续随时指出修改。

## 代码如何运行
```bash
cd server
npm install
npm run start:dev  # (自行添加package.json的脚本) 或者使用 npx ts-node src/main.ts
```