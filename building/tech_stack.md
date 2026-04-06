# MultiTrans 技术栈及基础设计文档
版本：1.0
日期：2026-04-06

## 1. 核心技术栈选择与评估

为达成“支持Windows/Linux/Android”、“高级现代前端”、“类Android Studio可视化参数调整”、“拖拽响应及剪贴板支持”和“系统原生分享能力”的综合需求，建议采用以下技术组合，均处于非常稳定且商业化友好的开源协议下。

| 组件 | 技术选项 | 推荐版本 | 许可证 | 官方文档 / 开发指南 |
| :--- | :--- | :--- | :--- | :--- |
| **服务器端 (Ubuntu)** | Node.js + NestJS | Node v20/v22+, NestJS v10+ | MIT | [NestJS Docs](https://docs.nestjs.com/) / [Node.js Docs](https://nodejs.org/docs/latest/) |
| **PC 客户端 (Win/Lin)** | Tauri + React + Tailwind CSS | Tauri v2, React v18+ | MIT / Apache 2.0 | [Tauri Docs](https://v2.tauri.app/) / [React Docs](https://react.dev/) / [TailwindCSS](https://tailwindcss.com/docs) |
| **Android 客户端** | Kotlin + Jetpack Compose | Kotlin 1.9+, Compose BOM 2024+ | Apache 2.0 | [Android Developers](https://developer.android.com/compose) |
| **客户端可视化构建** | Web: Plasmic / Builder.io / Figma<br>App: Compose Studio Preview | 最新版 | 商业/免费 | [Plasmic](https://www.plasmic.app/) / [Compose Tooling](https://developer.android.com/jetpack/compose/tooling) |
| **通信底层** | 局域网 WebSocket + REST | Socket.io v4+ | MIT | [Socket.io](https://socket.io/docs/v4/) |

---

## 2. 交互逻辑与实现指南

### 📱 移动端系统整合 (Android / OPPO 传送门兼容)
- **可视化设计:** 你提到“像使用 Android Studio 那样的 layout 设计界面”。如果采用 Jetpack Compose 构建 Android 端，你将继续享受 Android Studio 强大的 `@Preview` 注解系统。你可以利用参数调节来即时预览暗黑模式、大字号以及各类状态（完全符合可视化微调预期）。相比以前陈旧的 XML，它更现代且同样所见即所得。
- **发送方对接 (Share Intent):** 也就是 OPPO “发送给朋友” 以及提取文件的能力。在 `AndroidManifest.xml` 中需要声明接收隐式意图。即注册 `<intent-filter>` 接收 `android.intent.action.SEND` 和 `android.intent.action.SEND_MULTIPLE`，声明可以接收 `application/*` 或 `image/*` MIME 类型的分享内容。系统就会将应用暴露给文件分享入口。

### 💻 PC端交互体验 (跨端)
由于需要现代的 UI 以及强大的硬件交互（剪贴板、文件流），**Tauri 2.0** 是最佳基座，它使用网页语言（React/Vue）写前端，极大地降低开发难度并便于应用可视化工具：
- **拖拽文件与图片:** Tauri 提供原生的文件拖拽 API 与窗口监听事件（`tauri/api/window` -> `onFileDrop`）。
- **截取剪贴板内容:** 调用 Tauri 官方插件 `tauri-plugin-clipboard-manager` 读取文本和图片数据，摆脱传统浏览器的安全限制。

### 🤖 AI辅助开发与开源协议安全策略
- **版权声明隔离:** 为了避免“AI抄袭引发众矢之的”，使用开源基座时，优先借助该项目的官方文档与类型定义库 `node_modules/@types` 来驱动 AI。以上推荐的技术由于广泛采用最宽松的 MIT 或 Apache-2.0 协议，无需通过 GPL 传染，代码可以保留专有或许可证授权自由。
- **避免直接爬取代码段:** 以提供 API 文档（通过 DevDocs / MDN / Android 官方 GitHub）辅以 AI 构建出独有的业务逻辑，而非“搜索别人怎么写”去拼凑，这将带来实质性安全。

---

## 3. 下一步建议工作流

1. **环境准备**：使用 NVM 安装合适的 Node 环境，以及安装 Android Studio（集成 SDK）。对于 PC 端安装 Rust 环境用以编译 Tauri。
2. **初始化各端代码框架**：创建 `server`、`client-pc` 和 `client-android` 三个文件夹各自初始化。
3. **建立数据通信契约 (API Schema)**：比如传输图片是走 multipart/form-data 还是 Base64，是否需要建立连接密钥等。
