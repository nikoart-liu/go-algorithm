# Track Specification: Foundational Linear Data Structures

## Overview
本 Track 旨在为 Go 算法学习路径奠定坚实基础，实现最核心的线性数据结构。所有实现将采用 Go 泛型（Go 1.18+）以确保代码的通用性和类型安全，同时配备详尽的文档、图解和性能基准测试。

## Goals
1.  **实现核心结构**：单向链表 (Singly Linked List)、双向链表 (Doubly Linked List)、栈 (Stack)、队列 (Queue)、循环队列 (Circular Queue)。
2.  **工程化标准**：
    -   使用 Go 泛型。
    -   100% 单元测试覆盖（包含边界条件）。
    -   提供 `testing.B` 基准测试。
    -   通过 `golangci-lint` 检查。
3.  **可视化与文档**：
    -   每个结构包含 Mermaid.js 原理图。
    -   复杂度分析 (Big O)。
    -   性能对比报告。

## Detailed Requirements

### 1. Linked Lists (链表)
-   **Package**: `pkg/structures/list`
-   **Singly Linked List**:
    -   `Node[T]` 结构定义。
    -   操作：`PushFront`, `PushBack`, `InsertAt`, `Remove`, `Get`, `Size`, `Clear`.
-   **Doubly Linked List**:
    -   增加 `Prev` 指针。
    -   操作同上，优化反向遍历。
-   **Visuals**: 节点连接图，插入/删除操作流程图。

### 2. Stack (栈)
-   **Package**: `pkg/structures/stack`
-   **Implementations**:
    -   `SliceStack[T]`: 基于切片实现（高性能）。
    -   `ListStack[T]`: 基于链表实现（无扩容抖动）。
-   **Operations**: `Push`, `Pop`, `Peek`, `IsEmpty`, `Size`.
-   **Visuals**: LIFO (Last-In First-Out) 示意图。

### 3. Queue (队列)
-   **Package**: `pkg/structures/queue`
-   **Implementations**:
    -   `SliceQueue[T]`: 基于切片（注意缩容策略）。
    -   `ListQueue[T]`: 基于链表。
    -   `CircularQueue[T]`: 环形队列（数组实现，固定容量）。
-   **Operations**: `Enqueue`, `Dequeue`, `Peek`, `IsEmpty`, `Size`, `IsFull` (for Circular).
-   **Visuals**: FIFO (First-In First-Out) 示意图，环形指针移动图。

### 4. Performance & Documentation
-   **Benchmarks**:
    -   对比 Slice vs List 实现的性能 (Stack/Queue)。
    -   不同数据量级 (100, 10k, 1M) 下的操作耗时。
-   **README**:
    -   集成所有 Mermaid 图表。
    -   包含 Big O 表格。
    -   包含 Benchmark 结果分析。

## Tech Stack Alignment
-   **Language**: Go 1.21+
-   **Testing**: `testing`, `testify/assert`
-   **Linting**: `golangci-lint`
