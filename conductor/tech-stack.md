# Go 算法学习路径技术栈 (Tech Stack)

## 核心开发环境
- **编程语言**：Go 1.21+
- **版本管理**：Git

## 测试与性能分析工具
- **单元测试框架**：`testing` (标准库) + `github.com/stretchr/testify` (断言库)
- **基准测试**：`testing.B` (标准库)
- **性能分析**：`pprof` & `benchstat` (用于深度分析和对比测试数据)

## 文档与可视化
- **内容格式**：Markdown
- **图形渲染**：Mermaid.js (集成于 Markdown)
- **代码分析**：`golangci-lint` (保证代码质量和惯用法)

## 项目结构与依赖
- **包管理**：Go Modules
- **依赖策略**：**零外部依赖 (Zero External Dependencies)**。所有数据结构与算法均采用原生 Go 代码实现，以确保学习内容的完整性与纯粹性。
- **项目布局**：**标准 Go 布局 (Standard Go Layout)**。采用 `pkg/` 目录隔离具体实现，结构如下示例：
  - `/pkg/structures/...`: 数据结构（如链表、树、图）
  - `/pkg/algorithms/...`: 经典算法（如排序、搜索、动态规划）
  - `/cmd/...`: (可选) 用于演示或性能对比的 CLI 入口
